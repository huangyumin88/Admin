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
    <ElFormItem prop="user_id">
      <ElInputNumber v-model="queryCommon.data.user_id" :placeholder="t('banks.bankCards.name.user_id')" :min="1" :controls="false" />
    </ElFormItem>
    <ElFormItem prop="bank_id">
      <ElInputNumber v-model="queryCommon.data.bank_id" :placeholder="t('banks.bankCards.name.bank_id')" :min="1" :controls="false" />
    </ElFormItem>

<!--		<ElFormItem prop="user_id">-->
<!--			<MySelect v-model="queryCommon.data.user_id" :placeholder="t('banks.bankCards.name.user_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/user/user/list' }" />-->
<!--		</ElFormItem>-->
<!--		<ElFormItem prop="bank_id">-->
<!--			<MySelect v-model="queryCommon.data.bank_id" :placeholder="t('banks.bankCards.name.bank_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/banks/bank/list' }" />-->
<!--		</ElFormItem>-->
		<ElFormItem prop="card_number">
			<ElInput v-model="queryCommon.data.card_number" :placeholder="t('banks.bankCards.name.card_number')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="card_holder_name">
			<ElInput v-model="queryCommon.data.card_holder_name" :placeholder="t('banks.bankCards.name.card_holder_name')" minlength="1" maxlength="255" :clearable="true" />
		</ElFormItem>
		<ElFormItem prop="expiration_date">
			<ElDatePicker v-model="queryCommon.data.expiration_date" type="date" :placeholder="t('banks.bankCards.name.expiration_date')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
		</ElFormItem>
		<ElFormItem prop="cvv">
			<ElInput v-model="queryCommon.data.cvv" :placeholder="t('banks.bankCards.name.cvv')" minlength="1" maxlength="255" :clearable="true" />
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