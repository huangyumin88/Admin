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
		url: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
			{ type: 'url', trigger: 'change', message: t('validation.url') },
		],
		country_code: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/apple/cardUrl/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/apple/cardUrl/create', param, true)
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
				<ElFormItem :label="t('apple.cardUrl.name.url')" prop="url">
					<ElInput v-model="saveForm.data.url" :placeholder="t('apple.cardUrl.name.url')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('apple.cardUrl.name.country_code')" prop="country_code">
					<ElInput v-model="saveForm.data.country_code" :placeholder="t('apple.cardUrl.name.country_code')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('apple.cardUrl.name.isStop')" prop="isStop">
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