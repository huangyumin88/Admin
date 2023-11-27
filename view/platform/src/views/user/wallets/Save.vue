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
		user_id: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		],
		balance: [
			{ type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },	// 类型float值为0时验证不能通过
		],
		reward_points: [
			{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },
		],
		currency: [
			{ type: 'string', max: 10, trigger: 'blur', message: t('validation.max.string', { max: 10 }) },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/user/wallets/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/user/wallets/create', param, true)
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
<!--				<ElFormItem :label="t('user.wallets.name.user_id')" prop="user_id">-->
<!--					<MySelect v-model="saveForm.data.user_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/user/user/list' }" />-->
<!--				</ElFormItem>-->
				<ElFormItem :label="t('user.wallets.name.balance')" prop="balance">
					<ElInputNumber v-model="saveForm.data.balance" :placeholder="t('user.wallets.name.balance')" :precision="2" :controls="false"/>
				</ElFormItem>
				<ElFormItem :label="t('user.wallets.name.reward_points')" prop="reward_points">
					<ElInputNumber v-model="saveForm.data.reward_points" :placeholder="t('user.wallets.name.reward_points')" :min="0" :controls="false"/>
				</ElFormItem>
				<ElFormItem :label="t('user.wallets.name.currency')" prop="currency">
					<ElInput v-model="saveForm.data.currency" :placeholder="t('user.wallets.name.currency')" minlength="1" maxlength="10" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('user.wallets.name.isStop')" prop="isStop">
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