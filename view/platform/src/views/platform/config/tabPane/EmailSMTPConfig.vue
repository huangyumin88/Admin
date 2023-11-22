<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
  ref: null as any,
  loading: false,
  data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
    smtpType: 'smtpEmail',
    smtpHost: '',
    smtpPort: '',
    smtpEmail: '',
    smtpPwd: '',
    smtpTestEmail: '',
  } as { [propName: string]: any },
  rules: {

    smtpHost: [
      { type: 'string', trigger: 'blur', message: t('validation.input')},
    ],
    smtpPort: [
      { type: 'string', trigger: 'blur', message: t('validation.input') },
    ],
    smtpEmail: [
      { type: 'string', trigger: 'blur', message: t('validation.input') },
    ],
    smtpPwd: [
      { type: 'string', trigger: 'blur', message: t('validation.input') },
    ],
    smtpTestEmail: [
      { type: 'string', trigger: 'blur', message: t('validation.input') },
    ],
  } as any,
  initData: async () => {
    const param = { configKeyArr: Object.keys(saveForm.data) }
    try {
      const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
      saveForm.data = {
        ...saveForm.data,
        ...res.data.config
      }
    } catch (error) { }
  },
  submit: () => {
    saveForm.ref.validate(async (valid: boolean) => {
      if (!valid) {
        return false
      }
      saveForm.loading = true
      const param = removeEmptyOfObj(saveForm.data, false)
      try {
        await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/save', param, true)
      } catch (error) { }
      saveForm.loading = false
    })
  },
  test:async () => {
    saveForm.loading = true
    const param = removeEmptyOfObj(saveForm.data, false)
    try {
      await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/smtp/test', param, true)
    } catch (error) {
    }
    saveForm.loading = false
  },
  reset: () => {
    saveForm.ref.resetFields()
    saveForm.initData()
  }
})

saveForm.initData()
</script>

<template>
  <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto"
          :status-icon="true" :scroll-to-error="false">
<!--    <ElFormItem :label="t('platform.config.name.idCardType')" prop="idCardType">-->
<!--      <ElRadioGroup v-model="saveForm.data.idCardType">-->
<!--        <ElRadio v-for="(item, index) in (tm('platform.config.status.idCardType') as any)" :key="index"-->
<!--                 :label="item.value">-->
<!--          {{ item.label }}-->
<!--        </ElRadio>-->
<!--      </ElRadioGroup>-->
<!--    </ElFormItem>-->

    <template v-if="saveForm.data.smtpType == 'smtpEmail'">
      <ElFormItem :label="t('platform.config.name.emailSMTPHost')" prop="emailSMTPHost">
        <ElInput v-model="saveForm.data.smtpHost" :placeholder="t('platform.config.name.emailSMTPHost')"
                 :clearable="true" style="max-width: 500px;" />
        <label>
          <ElAlert type="info" :show-icon="true" :closable="false">
            <template #title>
<!--              <span v-html="t('platform.config.tip.aliyunIdCardHost')"></span>-->
              <span>
                Google Email 需要配置 SSL
              </span>
            </template>
          </ElAlert>
        </label>
      </ElFormItem>
      <ElFormItem :label="t('platform.config.name.emailSMTPPort')" prop="emailSMTPPort">
        <ElInput v-model="saveForm.data.smtpPort" :placeholder="t('platform.config.name.emailSMTPPort')"
                 :clearable="true" style="max-width: 500px;"/>
      </ElFormItem>
      <ElFormItem :label="t('platform.config.name.emailSMTPEmail')" prop="emailSMTPEmail">
        <ElInput v-model="saveForm.data.smtpEmail"
                 :placeholder="t('platform.config.name.emailSMTPEmail')" :clearable="true" style="max-width: 500px;"/>
      </ElFormItem>
      <ElFormItem :label="t('platform.config.name.emailSMTPPwd')" prop="emailSMTPPwd">
        <ElInput v-model="saveForm.data.smtpPwd"
                 :placeholder="t('platform.config.name.emailSMTPPwd')" :clearable="true" style="max-width: 500px;"/>
      </ElFormItem>
      <ElFormItem :label="t('platform.config.name.emailSMTPTestEmail')" prop="emailSMTPTestEmail">
        <ElInput v-model="saveForm.data.smtpTestEmail"
                 :placeholder="t('platform.config.name.emailSMTPTestEmail')" :clearable="true" style="max-width: 500px;"/>
      </ElFormItem>
    </template>

    <ElFormItem>
      <ElButton type="success" @click="saveForm.test" :loading="saveForm.loading">
        <AutoiconEpPromotion />{{ t('common.send') }}
      </ElButton>
      <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
        <AutoiconEpCircleCheck />{{ t('common.save') }}
      </ElButton>
      <ElButton type="info" @click="saveForm.reset">
        <AutoiconEpCircleClose />{{ t('common.reset') }}
      </ElButton>
    </ElFormItem>
  </ElForm>
</template>