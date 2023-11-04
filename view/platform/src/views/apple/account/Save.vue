<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		status: 0,
		login_status: 0,
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		account: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		pwd: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		country_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		balance: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
    stk: [
      { type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
    ],
		status: [
			{ type: 'enum', enum: (tm('apple.account.status.status') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
		],
		info: [],
		cookies: [],
		login_status: [
			{ type: 'enum', enum: (tm('apple.account.status.login_status') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
		],
		isStop: [
			{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
		],
	} as any,
	submit: () => {
		saveForm.ref.validate(async (valid: boolean) => {
			if (!valid) {
				return false
			}
			saveForm.loading = true
			const param = removeEmptyOfObj(saveForm.data, false)
			try {
				if (param?.idArr?.length > 0) {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/apple/account/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/apple/account/create', param, true)
				}
				listCommon.ref.getList(true)
				saveCommon.visible = false
			} catch (error) { }
			saveForm.loading = false
		})
	}
})

const saveDrawer = reactive({
	ref: null as any,
	size: useSettingStore().saveDrawer.size,
	beforeClose: (done: Function) => {
		if (useSettingStore().saveDrawer.isTipClose) {
			ElMessageBox.confirm('', {
				type: 'info',
				title: t('common.tip.configExit'),
				center: true,
				showClose: false,
			}).then(() => {
				done()
			}).catch(() => { })
		} else {
			done()
		}
	},
	buttonClose: () => {
		//saveCommon.visible = false
		saveDrawer.ref.handleClose()    //会触发beforeClose
	}
})
</script>

<template>
	<ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
		<ElScrollbar>
			<ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
				<ElFormItem :label="t('apple.account.name.account')" prop="account">
					<ElInput v-model="saveForm.data.account" :placeholder="t('apple.account.name.account')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.pwd')" prop="pwd">
					<ElInput v-model="saveForm.data.pwd" :placeholder="t('apple.account.name.pwd')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.country_id')" prop="country_id">
					<MySelect v-model="saveForm.data.country_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/apple/country/list' }" />
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.balance')" prop="balance">
					<ElInput v-model="saveForm.data.balance" :placeholder="t('apple.account.name.balance')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
        <ElFormItem :label="t('stk')" prop="stk">
          <ElInput v-model="saveForm.data.stk" :placeholder="t('stk')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
        </ElFormItem>
				<ElFormItem :label="t('apple.account.name.status')" prop="status">
					<ElRadioGroup v-model="saveForm.data.status">
						<ElRadio v-for="(item, index) in (tm('apple.account.status.status') as any)" :key="index" :label="item.value">
							{{ item.label }}
						</ElRadio>
					</ElRadioGroup>
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.info')" prop="info">
					<MyEditor v-model="saveForm.data.info" />
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.cookies')" prop="cookies">
					<MyEditor v-model="saveForm.data.cookies" />
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.login_status')" prop="login_status">
					<ElRadioGroup v-model="saveForm.data.login_status">
						<ElRadio v-for="(item, index) in (tm('apple.account.status.login_status') as any)" :key="index" :label="item.value">
							{{ item.label }}
						</ElRadio>
					</ElRadioGroup>
				</ElFormItem>
				<ElFormItem :label="t('apple.account.name.isStop')" prop="isStop">
					<ElSwitch v-model="saveForm.data.isStop" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
				</ElFormItem>
			</ElForm>
		</ElScrollbar>
		<template #footer>
			<ElButton @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</ElButton>
			<ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
				{{ t('common.save') }}
			</ElButton>
		</template>
	</ElDrawer>
</template>