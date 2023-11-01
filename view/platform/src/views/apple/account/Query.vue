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
	<ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
		<ElFormItem prop="id">
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
		</ElFormItem>
		<ElFormItem prop="account">
			<ElInput v-model="queryCommon.data.account" :placeholder="t('apple.account.name.account')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="pwd">
			<ElInput v-model="queryCommon.data.pwd" :placeholder="t('apple.account.name.pwd')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="country_id">
			<MySelect v-model="queryCommon.data.country_id" :placeholder="t('apple.account.name.country_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/apple/country/list' }" />
		</ElFormItem>
		<ElFormItem prop="balance">
			<ElInput v-model="queryCommon.data.balance" :placeholder="t('apple.account.name.balance')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="status" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.status" :options="tm('apple.account.status.status')" :placeholder="t('apple.account.name.status')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="login_status" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.login_status" :options="tm('apple.account.status.login_status')" :placeholder="t('apple.account.name.login_status')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isStop" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')" :placeholder="t('apple.account.name.isStop')" :clearable="true" />
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