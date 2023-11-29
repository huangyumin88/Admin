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
		// order_no: [
		// 	{ type: 'string', max: 255, trigger: 'blur', message: t('validation.max.string', { max: 255 }) },
		// ],
		// user_id: [
		// 	{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		// ],
		// salesperson_id: [
		// 	{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		// ],
		// client_status: [],
		// backend_status: [],
		// failed_reason: [],
		// failed_files: [],
		// trade_amount: [
		// 	{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },	// 类型float值为0时验证不能通过
		// ],
		// payable_amount: [
		// 	{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },
		// ],
		// card_cate_sub_id: [
		// 	{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
		// ],
		// device: [
		// 	{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) },
		// ],
		// wallet: [
		// 	{ type: 'string', min: 1, max: 10, trigger: 'blur', message: t('validation.max.string', { max: 10 }) },
		// ],
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
					await request(t('config.VITE_HTTP_API_PREFIX') + '/orders/orders/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/orders/orders/create', param, true)
				}
				listCommon.ref.getList(true)
				saveCommon.visible = false
			} catch (error) {
        //s
      }
			saveForm.loading = false
		})
	}
})

// 检查 queryCommon 是否存在并且是响应式的
if (saveForm && saveForm.data) {
  // 使用 watch 监听 card_cate_id 的变化
  watch(() => saveForm.data.card_cate_id, (newVal, oldVal) => {
    // console.log(`card_cate_id changed from ${oldVal} to ${newVal}`);
    // 在这里可以执行对应的逻辑
    saveForm.data.card_cate_sub_id = null;
  });
}

const isFormValid = computed(() => {
  return saveForm.data.card_cate_sub_id; // 示例验证逻辑
});

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
				<ElFormItem :label="t('orders.orders.name.order_no')" prop="order_no" v-if="saveCommon.title == '编辑'">
					<ElInput v-model="saveForm.data.order_no" :placeholder="t('orders.orders.name.order_no')" disabled minlength="1" maxlength="255" :show-word-limit="true" :clearable="true" />
				</ElFormItem>
				<ElFormItem :label="t('orders.orders.name.user_id')" prop="user_id" v-if="saveCommon.title == '新增'">
					<MySelect  v-model="saveForm.data.user_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/user/user/list' }" />
				</ElFormItem>

<!--        <ElFormItem :label="t('orders.orders.name.salesperson_id')" prop="salesperson_id">-->
<!--					<MySelect v-model="saveForm.data.salesperson_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/salesperson/list' }" />-->
<!--				</ElFormItem>-->
<!--				<ElFormItem :label="t('orders.orders.name.client_status')" prop="client_status">-->
<!--					<ElInput v-model="saveForm.data.client_status" :placeholder="t('orders.orders.name.client_status')" disabled :show-word-limit="true" :clearable="true" />-->
<!--				</ElFormItem>-->
				<ElFormItem :label="t('orders.orders.name.backend_status')" prop="backend_status">
          <MySelect v-model="saveForm.data.backend_status" :placeholder="t('orders.orders.name.backend_status')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/orders/orders/checkorderstatus',param:{status: saveCommon.data.backend_status} }" />
<!--					<ElInput v-model="saveForm.data.backend_status" :placeholder="t('orders.orders.name.backend_status')" :show-word-limit="true" :clearable="true" />-->
				</ElFormItem>
<!--				<ElFormItem :label="t('orders.orders.name.failed_reason')" prop="failed_reason">-->
<!--					<MyEditor v-model="saveForm.data.failed_reason" />-->
<!--				</ElFormItem>-->


        <ElFormItem :label="t('orders.orders.name.failed_reason')" prop="failed_reason" v-if="saveForm.data.backend_status == 'Failed'">
          <ElInput v-model="saveForm.data.failed_reason" :placeholder="t('orders.orders.name.failed_reason')" minlength="1" maxlength="1000" :show-word-limit="true" :clearable="true" :autosize="{ minRows: 6, maxRows: 8 }"
                   type="textarea"/>
        </ElFormItem>


				<ElFormItem :label="t('orders.orders.name.failed_files')" prop="failed_files" v-if="saveForm.data.backend_status == 'Failed'">
<!--					<MyEditor v-model="saveForm.data.failed_files" />-->
          <MyUpload v-model="saveForm.data.failed_files" accept="image/*" />
				</ElFormItem>

        <ElFormItem :label="t('orders.orders.name.card_cate_id')" prop="card_cate_id">
          <MySelect v-model="saveForm.data.card_cate_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/cardCategories/list' }" style="width: 350px"/>
        </ElFormItem>

				<ElFormItem :label="t('orders.orders.name.card_cate_sub_id')" prop="card_cate_sub_id" v-if="saveForm.data.card_cate_id">
					<MySelect v-model="saveForm.data.card_cate_sub_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/cardCategoriesSub/list',param: { filter: {cate_id:saveForm.data.card_cate_id}, field:['id','label']}}" style="width: 350px" />
				</ElFormItem>

        <ElFormItem :label="t('orders.orders.name.trade_amount')" prop="trade_amount">
          <ElInputNumber v-model="saveForm.data.trade_amount" :placeholder="t('orders.orders.name.trade_amount')" :min="0"  :controls="false" :disabled="!isFormValid" style="margin-right: 10px;" />
          <label>
            <ElAlert type="info" :show-icon="true" :closable="false">
              <template #title>
                <span>
                请先选择兑换卡
              </span>
              </template>
            </ElAlert>
          </label>

        </ElFormItem>

        <ElFormItem :label="t('orders.orders.name.payable_amount')" prop="payable_amount">
          <ElInputNumber v-model="saveForm.data.payable_amount" :placeholder="t('orders.orders.name.payable_amount')" :min="0" :controls="false" :disabled="!isFormValid"/>

          <label>
            <ElAlert type="info" :show-icon="true" :closable="false">
              <template #title>
                <span>
                根据兑换卡子分类的汇率 填写NGN 支付金额
              </span>
              </template>
            </ElAlert>
          </label>

        </ElFormItem>

        <ElFormItem :label="t('orders.orders.name.trade_files')" prop="trade_files">
<!--          <MyEditor v-model="saveForm.data.trade_files" />-->
          <MyUpload v-model="saveForm.data.trade_files" accept="image/*" />
        </ElFormItem>

<!--				<ElFormItem :label="t('orders.orders.name.device')" prop="device">-->
<!--					<ElInput v-model="saveForm.data.device" :placeholder="t('orders.orders.name.device')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" />-->
<!--				</ElFormItem>-->
<!--				<ElFormItem :label="t('orders.orders.name.wallet')" prop="wallet">-->
<!--					<ElInput v-model="saveForm.data.wallet" :placeholder="t('orders.orders.name.wallet')" minlength="1" maxlength="10" :show-word-limit="true" :clearable="true" />-->
<!--				</ElFormItem>-->
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