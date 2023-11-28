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
		<ElFormItem prop="order_no">
			<ElInput v-model="queryCommon.data.order_no" :placeholder="t('orders.orders.name.order_no')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="user_id">
			<MySelect v-model="queryCommon.data.user_id" :placeholder="t('orders.orders.name.user_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/user/user/list' }" />
		</ElFormItem>
		<ElFormItem prop="salesperson_id">
			<MySelect v-model="queryCommon.data.salesperson_id" :placeholder="t('orders.orders.name.salesperson_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/salesperson/list' }" />
		</ElFormItem>
		<ElFormItem prop="client_status">
			<ElInput v-model="queryCommon.data.client_status" :placeholder="t('orders.orders.name.client_status')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="backend_status">
			<ElInput v-model="queryCommon.data.backend_status" :placeholder="t('orders.orders.name.backend_status')" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="card_cate_sub_id">
			<MySelect v-model="queryCommon.data.card_cate_sub_id" :placeholder="t('orders.orders.name.card_cate_sub_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/cardCateSub/list' }" />
		</ElFormItem>
		<ElFormItem prop="device">
			<ElInput v-model="queryCommon.data.device" :placeholder="t('orders.orders.name.device')" minlength="1" maxlength="30" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="wallet">
			<ElInput v-model="queryCommon.data.wallet" :placeholder="t('orders.orders.name.wallet')" minlength="1" maxlength="10" :clearable="true" />
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