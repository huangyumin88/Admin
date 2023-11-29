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
		id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		actions_user_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		order_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		backend_status: [
			{ type: 'string', max: 10, trigger: 'blur', message: t('validation.max.string', { max: 10 }) },
		],
		remarks: [
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/orders/ordersActions/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/orders/ordersActions/create', param, true)
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
				<ElFormItem :label="t('orders.ordersActions.name.id')" prop="id">
					<MySelect v-model="saveForm.data.id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders//list' }" />
				</ElFormItem>
				<ElFormItem :label="t('orders.ordersActions.name.actions_user_id')" prop="actions_user_id">
					<MySelect v-model="saveForm.data.actions_user_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/actionsUser/list' }" />
				</ElFormItem>
				<ElFormItem :label="t('orders.ordersActions.name.order_id')" prop="order_id">
					<MySelect v-model="saveForm.data.order_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/order/list' }" />
				</ElFormItem>
				<ElFormItem :label="t('orders.ordersActions.name.backend_status')" prop="backend_status">
					<ElInput v-model="saveForm.data.backend_status" :placeholder="t('orders.ordersActions.name.backend_status')" minlength="1" maxlength="10" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('orders.ordersActions.name.remarks')" prop="remarks">
					<ElInput v-model="saveForm.data.remarks" :placeholder="t('orders.ordersActions.name.remarks')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
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