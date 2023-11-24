<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		sort: 50,
		isActive: 1,
		isStop: 0,
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {
		sub_id: [
			{ type: 'string', max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
		],
		avatar: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
		],
		avatar_url: [
			{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ type: 'url', trigger: 'change', message: t('validation.url') },
		],
		name: [
			{ type: 'string', required: true, max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
		],
		sort: [
			{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
		],
		isActive: [
			{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/app/cardCategories/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/app/cardCategories/create', param, true)
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
				<ElFormItem :label="t('app.cardCategories.name.sub_id')" prop="sub_id">
					<ElInput v-model="saveForm.data.sub_id" :placeholder="t('app.cardCategories.name.sub_id')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.avatar')" prop="avatar">
					<MyUpload v-model="saveForm.data.avatar" accept="image/*" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.avatar_url')" prop="avatar_url">
					<ElInput v-model="saveForm.data.avatar_url" :placeholder="t('app.cardCategories.name.avatar_url')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.name')" prop="name">
					<ElInput v-model="saveForm.data.name" :placeholder="t('app.cardCategories.name.name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.sort')" prop="sort">
					<ElInputNumber v-model="saveForm.data.sort" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="50" />
					<label>
						<ElAlert :title="t('app.cardCategories.tip.sort')" type="info" :show-icon="true" :closable="false" />
					</label>
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.isActive')" prop="isActive">
					<ElSwitch v-model="saveForm.data.isActive" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
				</ElFormItem>
				<ElFormItem :label="t('app.cardCategories.name.isStop')" prop="isStop">
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