<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		isStop: 0,
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		name: [
			{ type: 'string', required: true, max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		slug: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
		],
		code: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		flutterwaveBankCode: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		redbillerBankCode: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		anchorBankCode: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		country: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
		],
		currency: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
		],
		type: [
			{ type: 'string', max: 50, trigger: 'blur', message: t('validation.max.string', { max: 50 }) },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/banks/banks/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/banks/banks/create', param, true)
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
		saveDrawer.ref.handleClose()	//会触发beforeClose
	}
})
</script>

<template>
	<ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
		<ElScrollbar>
			<ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
				<ElFormItem :label="t('banks.banks.name.name')" prop="name">
					<ElInput v-model="saveForm.data.name" :placeholder="t('banks.banks.name.name')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.slug')" prop="slug">
					<ElInput v-model="saveForm.data.slug" :placeholder="t('banks.banks.name.slug')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.code')" prop="code">
					<ElInput v-model="saveForm.data.code" :placeholder="t('banks.banks.name.code')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.flutterwaveBankCode')" prop="flutterwaveBankCode">
					<ElInput v-model="saveForm.data.flutterwaveBankCode" :placeholder="t('banks.banks.name.flutterwaveBankCode')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.redbillerBankCode')" prop="redbillerBankCode">
					<ElInput v-model="saveForm.data.redbillerBankCode" :placeholder="t('banks.banks.name.redbillerBankCode')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.anchorBankCode')" prop="anchorBankCode">
					<ElInput v-model="saveForm.data.anchorBankCode" :placeholder="t('banks.banks.name.anchorBankCode')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.country')" prop="country">
					<ElInput v-model="saveForm.data.country" :placeholder="t('banks.banks.name.country')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.currency')" prop="currency">
					<ElInput v-model="saveForm.data.currency" :placeholder="t('banks.banks.name.currency')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.type')" prop="type">
					<ElInput v-model="saveForm.data.type" :placeholder="t('banks.banks.name.type')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.banks.name.isStop')" prop="isStop">
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