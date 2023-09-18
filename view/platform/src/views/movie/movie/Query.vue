<script setup lang="ts">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
	...queryCommon.data,
	timeRange: (() => {
		// const date = new Date()
		return [
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
		]
	})(),
	timeRangeStart: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),
	timeRangeEnd: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
	ref: null as any,
	loading: false,
	submit: () => {
		queryForm.loading = true
		listCommon.ref.getList(true).finally(() => {
			queryForm.loading = false
		})
	},
	reset: () => {
		queryForm.ref.resetFields()
		//queryForm.submit()
	}
})
</script>

<template>
	<ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryCommon.data" :inline="true"
		@keyup.enter="queryForm.submit">
		<ElFormItem prop="id">
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
		</ElFormItem>
		<ElFormItem prop="movie_id">
			<ElInput v-model="queryCommon.data.movie_id" :placeholder="t('movie.movie.name.movie_id')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="movie_name">
			<ElInput v-model="queryCommon.data.movie_name" :placeholder="t('movie.movie.name.movie_name')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="movie_real_name">
			<ElInput v-model="queryCommon.data.movie_real_name" :placeholder="t('movie.movie.name.movie_real_name')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="movie_picture">
			<ElInput v-model="queryCommon.data.movie_picture" :placeholder="t('movie.movie.name.movie_picture')" minlength="1" maxlength="500" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="plot_photo">
			<ElInput v-model="queryCommon.data.plot_photo" :placeholder="t('movie.movie.name.plot_photo')" minlength="1" maxlength="500" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="translated_name">
			<ElInput v-model="queryCommon.data.translated_name" :placeholder="t('movie.movie.name.translated_name')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="url">
			<ElInput v-model="queryCommon.data.url" :placeholder="t('movie.movie.name.url')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="years">
			<ElInput v-model="queryCommon.data.years" :placeholder="t('movie.movie.name.years')" minlength="1" maxlength="25" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="place">
			<ElInput v-model="queryCommon.data.place" :placeholder="t('movie.movie.name.place')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="category_name">
			<ElInput v-model="queryCommon.data.category_name" :placeholder="t('movie.movie.name.category_name')" minlength="1" maxlength="50" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="category_ids">
			<ElInput v-model="queryCommon.data.category_ids" :placeholder="t('movie.movie.name.category_ids')" minlength="1" maxlength="50" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="language_name">
			<ElInput v-model="queryCommon.data.language_name" :placeholder="t('movie.movie.name.language_name')" minlength="1" maxlength="100" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="subtitle">
			<ElInput v-model="queryCommon.data.subtitle" :placeholder="t('movie.movie.name.subtitle')" minlength="1" maxlength="25" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="release_date">
			<ElInput v-model="queryCommon.data.release_date" :placeholder="t('movie.movie.name.release_date')" minlength="1" maxlength="100" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="imdb">
			<ElInput v-model="queryCommon.data.imdb" :placeholder="t('movie.movie.name.imdb')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="doubandb">
			<ElInput v-model="queryCommon.data.doubandb" :placeholder="t('movie.movie.name.doubandb')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="score">
			<ElInput v-model="queryCommon.data.score" :placeholder="t('movie.movie.name.score')" minlength="1" maxlength="5" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="awards">
			<ElInput v-model="queryCommon.data.awards" :placeholder="t('movie.movie.name.awards')" minlength="1" maxlength="500" :clearable="true" />
		</ElFormItem>
<!--		<ElFormItem prop="times">-->
<!--			<ElInput v-model="queryCommon.data.times" :placeholder="t('movie.movie.name.times')" minlength="1" maxlength="500" :clearable="true" />-->
<!--		</ElFormItem>-->
<!--		<ElFormItem prop="time">-->
<!--			<ElInput v-model="queryCommon.data.time" :placeholder="t('movie.movie.name.time')" minlength="1" maxlength="25" :clearable="true" />-->
<!--		</ElFormItem>-->
		<ElFormItem prop="director">
			<ElInput v-model="queryCommon.data.director" :placeholder="t('movie.movie.name.director')" minlength="1" maxlength="50" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="scriptwriter">
			<ElInput v-model="queryCommon.data.scriptwriter" :placeholder="t('movie.movie.name.scriptwriter')" minlength="1" maxlength="100" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="video_size">
			<ElInput v-model="queryCommon.data.video_size" :placeholder="t('movie.movie.name.video_size')" minlength="1" maxlength="25" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="video_resolution">
			<ElInput v-model="queryCommon.data.video_resolution" :placeholder="t('movie.movie.name.video_resolution')" minlength="1" maxlength="25" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isUpdate" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.isUpdate" :options="tm('common.status.whether')" :placeholder="t('movie.movie.name.isUpdate')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isStop" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('movie.movie.name.isStop')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="timeRange">
			<ElDatePicker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
		</ElFormItem>
		<ElFormItem>
			<ElButton type="primary" @click="queryForm.submit" :loading="queryForm.loading">
				<AutoiconEpSearch />{{ t('common.query') }}
			</ElButton>
			<ElButton type="info" @click="queryForm.reset">
				<AutoiconEpCircleClose />{{ t('common.reset') }}
			</ElButton>
		</ElFormItem>
	</ElForm>
</template>