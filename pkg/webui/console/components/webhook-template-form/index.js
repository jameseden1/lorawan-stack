// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { Component } from 'react'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'
import urlTemplate from 'url-template'

import tts from '@console/api/tts'

import Form from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'
import PortalledModal from '@ttn-lw/components/modal/portalled'

import WebhookTemplateInfo from '@console/components/webhook-template-info'

import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import { id as webhookIdRegexp } from '@ttn-lw/lib/regexp'

const m = defineMessages({
  createTemplate: 'Create {template} webhook',
  idPlaceholder: 'my-new-{templateId}-webhook',
  templateSettings: 'Template settings',
})

const pathExpand = (url, fields) =>
  Boolean(url) && url.path ? { path: urlTemplate.parse(url.path).expand(fields) } : url

export default class WebhookTemplateForm extends Component {
  static propTypes = {
    appId: PropTypes.string.isRequired,
    existCheck: PropTypes.func.isRequired,
    onSubmit: PropTypes.func.isRequired,
    onSubmitSuccess: PropTypes.func.isRequired,
    templateId: PropTypes.string.isRequired,
    webhookTemplate: PropTypes.webhookTemplate.isRequired,
  }

  modalResolve = () => null
  modalReject = () => null

  state = {
    error: undefined,
    displayOverwriteModal: false,
    existingId: undefined,
  }

  @bind
  async convertTemplateToWebhook(values) {
    const { webhookTemplate: template, appId } = this.props
    const { webhook_id, ...fields } = values

    const headers = Object.keys(template.headers || {}).reduce((acc, key) => {
      const val = template.headers[key]
      const headerValue = val.replace(/{([a-z0-9_-]+)}/i, (_, field) => values[field])
      if (headerValue !== '') {
        acc[key] = headerValue
      }
      return acc
    }, {})

    const webhook = {
      ids: {
        webhook_id: values.webhook_id,
      },
      template_ids: template.ids,
      format: template.format,
      headers,
      template_fields: fields,
      base_url: urlTemplate.parse(template.base_url).expand(fields),
      uplink_message: pathExpand(template.uplink_message, fields),
      join_accept: pathExpand(template.join_accept, fields),
      downlink_ack: pathExpand(template.downlink_ack, fields),
      downlink_nack: pathExpand(template.downlink_nack, fields),
      downlink_sent: pathExpand(template.downlink_sent, fields),
      downlink_failed: pathExpand(template.downlink_failed, fields),
      downlink_queued: pathExpand(template.downlink_queued, fields),
      downlink_queue_invalidated: pathExpand(template.downlink_queue_invalidated, fields),
      location_solved: pathExpand(template.location_solved, fields),
      service_data: pathExpand(template.service_data, fields),
    }

    if (template.create_downlink_api_key) {
      const key = {
        name: `${webhook_id} downlink API key`,
        rights: ['RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE'],
      }
      const { key: downlink_api_key } = await tts.Applications.ApiKeys.create(appId, key)
      webhook.downlink_api_key = downlink_api_key
    }

    return webhook
  }

  @bind
  async handleSubmit(values, { setSubmitting, resetForm }) {
    const { onSubmit, onSubmitSuccess, existCheck } = this.props

    await this.setState({ error: '' })
    try {
      const webhook = await this.convertTemplateToWebhook(values)

      const webhookId = webhook.ids.webhook_id
      const exists = await existCheck(webhookId)
      if (exists) {
        this.setState({ displayOverwriteModal: true, existingId: webhookId })
        await new Promise((resolve, reject) => {
          this.modalResolve = resolve
          this.modalReject = reject
        })
      }

      const result = await onSubmit(webhook)
      resetForm({ values })
      onSubmitSuccess(result)
    } catch (error) {
      setSubmitting(false)
      await this.setState({ error })
    }
  }

  @bind
  handleReplaceModalDecision(mayReplace) {
    if (mayReplace) {
      this.modalResolve()
    } else {
      this.modalReject()
    }
    this.setState({ displayOverwriteModal: false })
  }

  render() {
    const { templateId, webhookTemplate } = this.props
    const { name, fields = [] } = webhookTemplate
    const { error, displayOverwriteModal, existingId } = this.state
    const validationSchema = Yup.object({
      ...fields.reduce(
        (acc, field) => ({
          ...acc,
          [field.id]: field.optional
            ? Yup.string()
            : Yup.string().required(sharedMessages.validateRequired),
        }),
        {
          webhook_id: Yup.string()
            .min(3, Yup.passValues(sharedMessages.validateTooShort))
            .max(36, Yup.passValues(sharedMessages.validateTooLong))
            .matches(webhookIdRegexp, Yup.passValues(sharedMessages.validateIdFormat))
            .required(sharedMessages.validateRequired),
        },
      ),
    })

    const initialValues = fields.reduce(
      (acc, field) => ({ ...acc, [field.id]: field.default_value || '' }),
      {
        webhook_id: '',
      },
    )
    return (
      <div>
        <PortalledModal
          title={sharedMessages.idAlreadyExists}
          message={{
            ...sharedMessages.webhookAlreadyExistsModalMessage,
            values: { id: existingId },
          }}
          buttonMessage={sharedMessages.replaceWebhook}
          onComplete={this.handleReplaceModalDecision}
          approval
          visible={displayOverwriteModal}
        />
        <WebhookTemplateInfo webhookTemplate={webhookTemplate} />
        <Form
          onSubmit={this.handleSubmit}
          validationSchema={validationSchema}
          initialValues={initialValues}
          error={error}
          formikRef={this.form}
        >
          <Form.SubTitle title={m.templateSettings} />
          <Form.Field
            name="webhook_id"
            title={sharedMessages.webhookId}
            placeholder={{ ...m.idPlaceholder, values: { templateId } }}
            component={Input}
            required
            autoFocus
          />
          {fields.map(field => (
            <Form.Field
              component={Input}
              name={field.id}
              title={field.name}
              description={field.description}
              key={field.id}
              required={!field.optional}
              sensitive={field.secret}
            />
          ))}
          <SubmitBar>
            <Form.Submit
              component={SubmitButton}
              message={{
                ...m.createTemplate,
                values: { template: name.toLowerCase() },
              }}
            />
          </SubmitBar>
        </Form>
      </div>
    )
  }
}
