<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		user_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		bank_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		card_number: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
		],
		card_holder_name: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		expiration_date: [
			{ type: 'string', trigger: 'change', message: t('validation.select') },
		],
		cvv: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/banks/bankCards/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/banks/bankCards/create', param, true)
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
				<ElFormItem :label="t('banks.bankCards.name.user_id')" prop="user_id">
					<MySelect v-model="saveForm.data.user_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/user/user/list' }" />
				</ElFormItem>
				<ElFormItem :label="t('banks.bankCards.name.bank_id')" prop="bank_id">
					<MySelect v-model="saveForm.data.bank_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/banks/bank/list' }" />
				</ElFormItem>
				<ElFormItem :label="t('banks.bankCards.name.card_number')" prop="card_number">
					<ElInput v-model="saveForm.data.card_number" :placeholder="t('banks.bankCards.name.card_number')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.bankCards.name.card_holder_name')" prop="card_holder_name">
					<ElInput v-model="saveForm.data.card_holder_name" :placeholder="t('banks.bankCards.name.card_holder_name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('banks.bankCards.name.expiration_date')" prop="expiration_date">
					<ElDatePicker v-model="saveForm.data.expiration_date" type="date" :placeholder="t('banks.bankCards.name.expiration_date')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
				</ElFormItem>
				<ElFormItem :label="t('banks.bankCards.name.cvv')" prop="cvv">
					<ElInput v-model="saveForm.data.cvv" :placeholder="t('banks.bankCards.name.cvv')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
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