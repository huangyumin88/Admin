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
			{ type: 'string', required: true, max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		isoName: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		currencyCode: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		currencyName: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		flagUrl: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ type: 'url', trigger: 'change', message: t('validation.url') },
		],
		flagAvatar: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
		],
		flagAvatarID: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/app/cardCountries/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/app/cardCountries/create', param, true)
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
				<ElFormItem :label="t('app.cardCountries.name.name')" prop="name">
					<ElInput v-model="saveForm.data.name" :placeholder="t('app.cardCountries.name.name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.isoName')" prop="isoName">
					<ElInput v-model="saveForm.data.isoName" :placeholder="t('app.cardCountries.name.isoName')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.currencyCode')" prop="currencyCode">
					<ElInput v-model="saveForm.data.currencyCode" :placeholder="t('app.cardCountries.name.currencyCode')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.currencyName')" prop="currencyName">
					<ElInput v-model="saveForm.data.currencyName" :placeholder="t('app.cardCountries.name.currencyName')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.flagUrl')" prop="flagUrl">
					<ElInput v-model="saveForm.data.flagUrl" :placeholder="t('app.cardCountries.name.flagUrl')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.flagAvatar')" prop="flagAvatar">
					<ElInput v-model="saveForm.data.flagAvatar" :placeholder="t('app.cardCountries.name.flagAvatar')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.flagAvatarID')" prop="flagAvatarID">
					<ElInput v-model="saveForm.data.flagAvatarID" :placeholder="t('app.cardCountries.name.flagAvatarID')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCountries.name.isStop')" prop="isStop">
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