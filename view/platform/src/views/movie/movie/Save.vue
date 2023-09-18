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
		movie_id: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		movie_name: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		movie_real_name: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		movie_picture: [
			{ type: 'string', min: 1, max: 500, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 500 }) },
		],
		plot_photo: [
			{ type: 'string', min: 1, max: 500, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 500 }) },
		],
		translated_name: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		sub_intros: [],
		url: [
			{ type: 'url', trigger: 'change', message: t('validation.url') },
			{ type: 'string', min: 1, max: 255, trigger: 'change', message: t('validation.between.string', { min: 1, max: 255 }) }
		],
		years: [
			{ type: 'string', min: 1, max: 25, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 25 }) },
		],
		place: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		category_name: [
			{ type: 'string', min: 1, max: 50, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 50 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		category_ids: [
			{ type: 'string', min: 1, max: 50, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 50 }) },
		],
		language_name: [
			{ type: 'string', min: 1, max: 100, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 100 }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],
		subtitle: [
			{ type: 'string', min: 1, max: 25, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 25 }) },
		],
		release_date: [
			{ type: 'string', min: 1, max: 100, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 100 }) },
		],
		imdb: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		doubandb: [
			{ type: 'string', min: 1, max: 255, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 255 }) },
		],
		score: [
			{ type: 'string', min: 1, max: 5, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 5 }) },
		],
		awards: [
			{ type: 'string', min: 1, max: 500, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 500 }) },
		],
		times: [
			{ type: 'string', min: 1, max: 500, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 500 }) },
		],
		time: [
			{ type: 'string', min: 1, max: 25, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 25 }) },
		],
		director: [
			{ type: 'string', min: 1, max: 50, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 50 }) },
		],
		scriptwriter: [
			{ type: 'string', min: 1, max: 100, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 100 }) },
		],
		actors: [],
		intros: [],
		down_urls: [],
		video_size: [
			{ type: 'string', min: 1, max: 25, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 25 }) },
		],
		video_resolution: [
			{ type: 'string', min: 1, max: 25, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 25 }) },
		],
		isUpdate: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],
		isStop: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/movie/movie/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/movie/movie/create', param, true)
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
	<ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible"
		:title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
		<ElScrollbar>
			<ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
				label-width="auto" :status-icon="true" :scroll-to-error="true">
				<ElFormItem :label="t('movie.movie.name.movie_id')" prop="movie_id">
					<ElInput v-model="saveForm.data.movie_id" :placeholder="t('movie.movie.name.movie_id')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.movie_name')" prop="movie_name">
					<ElInput v-model="saveForm.data.movie_name" :placeholder="t('movie.movie.name.movie_name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.movie_real_name')" prop="movie_real_name">
					<ElInput v-model="saveForm.data.movie_real_name" :placeholder="t('movie.movie.name.movie_real_name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.movie_picture')" prop="movie_picture">
					<ElInput v-model="saveForm.data.movie_picture" :placeholder="t('movie.movie.name.movie_picture')" minlength="1" maxlength="500" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.plot_photo')" prop="plot_photo">
					<ElInput v-model="saveForm.data.plot_photo" :placeholder="t('movie.movie.name.plot_photo')" minlength="1" maxlength="500" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.translated_name')" prop="translated_name">
					<ElInput v-model="saveForm.data.translated_name" :placeholder="t('movie.movie.name.translated_name')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.sub_intros')" prop="sub_intros">
					<MyEditor v-model="saveForm.data.sub_intros" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.url')" prop="url">
					<ElInput v-model="saveForm.data.url" :placeholder="t('movie.movie.name.url')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.years')" prop="years">
					<ElInput v-model="saveForm.data.years" :placeholder="t('movie.movie.name.years')" minlength="1" maxlength="25" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.place')" prop="place">
					<ElInput v-model="saveForm.data.place" :placeholder="t('movie.movie.name.place')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.category_name')" prop="category_name">
					<ElInput v-model="saveForm.data.category_name" :placeholder="t('movie.movie.name.category_name')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.category_ids')" prop="category_ids">
					<ElInput v-model="saveForm.data.category_ids" :placeholder="t('movie.movie.name.category_ids')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.language_name')" prop="language_name">
					<ElInput v-model="saveForm.data.language_name" :placeholder="t('movie.movie.name.language_name')" minlength="1" maxlength="100" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.subtitle')" prop="subtitle">
					<ElInput v-model="saveForm.data.subtitle" :placeholder="t('movie.movie.name.subtitle')" minlength="1" maxlength="25" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.release_date')" prop="release_date">
					<ElInput v-model="saveForm.data.release_date" :placeholder="t('movie.movie.name.release_date')" minlength="1" maxlength="100" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.imdb')" prop="imdb">
					<ElInput v-model="saveForm.data.imdb" :placeholder="t('movie.movie.name.imdb')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.doubandb')" prop="doubandb">
					<ElInput v-model="saveForm.data.doubandb" :placeholder="t('movie.movie.name.doubandb')" minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.score')" prop="score">
					<ElInput v-model="saveForm.data.score" :placeholder="t('movie.movie.name.score')" minlength="1" maxlength="5" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.awards')" prop="awards">
					<ElInput v-model="saveForm.data.awards" :placeholder="t('movie.movie.name.awards')" minlength="1" maxlength="500" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.times')" prop="times">
					<ElInput v-model="saveForm.data.times" :placeholder="t('movie.movie.name.times')" minlength="1" maxlength="500" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.time')" prop="time">
					<ElInput v-model="saveForm.data.time" :placeholder="t('movie.movie.name.time')" minlength="1" maxlength="25" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.director')" prop="director">
					<ElInput v-model="saveForm.data.director" :placeholder="t('movie.movie.name.director')" minlength="1" maxlength="50" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.scriptwriter')" prop="scriptwriter">
					<ElInput v-model="saveForm.data.scriptwriter" :placeholder="t('movie.movie.name.scriptwriter')" minlength="1" maxlength="100" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.actors')" prop="actors">
					<MyEditor v-model="saveForm.data.actors" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.intros')" prop="intros">
					<MyEditor v-model="saveForm.data.intros" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.down_urls')" prop="down_urls">
					<MyEditor v-model="saveForm.data.down_urls" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.video_size')" prop="video_size">
					<ElInput v-model="saveForm.data.video_size" :placeholder="t('movie.movie.name.video_size')" minlength="1" maxlength="25" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.video_resolution')" prop="video_resolution">
					<ElInput v-model="saveForm.data.video_resolution" :placeholder="t('movie.movie.name.video_resolution')" minlength="1" maxlength="25" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('movie.movie.name.isUpdate')" prop="isUpdate">
                    <ElSwitch v-model="saveForm.data.isUpdate" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </ElFormItem>
				<ElFormItem :label="t('movie.movie.name.isStop')" prop="isStop">
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