<script setup lang="tsx">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
        sceneId: saveCommon.data.sceneId ? saveCommon.data.sceneId : undefined,
        // tableId: saveCommon.data.tableId ? saveCommon.data.tableId : undefined,
    } as { [propName: string]: any },
    rules: {
        roleName: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
        ],
        sceneId: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 1, message: t('validation.select') },
        ],
        /* tableId: [
            // { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 0, message: t('validation.select') },
        ], */
        isStop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        menuIdArr: [
            {
                type: 'array',
                trigger: 'change',
                message: t('validation.select'),
                defaultField: {
                    type: 'array',
                    defaultField: {
                        type: 'integer',
                        min: 1,
                        message: t('validation.min.number', { min: 1 }),
                    },
                },
            },
        ],
        actionIdArr: [{ type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'integer', min: 1, message: t('validation.min.number', { min: 1 }) } }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        console.log(saveForm.data.menuIdArr)
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.sceneId === undefined ? (param.sceneId = 0) : null
            // param.tableId === undefined ? (param.tableId = 0) : null
            if (param.menuIdArr === undefined) {
                param.menuIdArr = []
            } else {
                let menuIdArr: any = []
                param.menuIdArr.forEach((item: any) => {
                    menuIdArr = menuIdArr.concat(item)
                })
                //param.menuIdArr = [...new Set(menuIdArr)]
                param.menuIdArr = menuIdArr.filter((item: any, index: any) => {
                    return menuIdArr.indexOf(item) === index
                })
            }
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/role/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/auth/role/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
        })
    },
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
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('auth.role.name.roleName')" prop="roleName">
                    <el-input v-model="saveForm.data.roleName" :placeholder="t('auth.role.name.roleName')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('auth.role.name.sceneId')" prop="sceneId">
                    <my-select
                        v-model="saveForm.data.sceneId"
                        :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }"
                        @change="
                            () => {
                                saveForm.data.menuIdArr = []
                                saveForm.data.actionIdArr = []
                            }
                        "
                    />
                </el-form-item>
                <!-- <el-form-item :label="t('auth.role.name.tableId')" prop="tableId">
                    <my-select v-model="saveForm.data.tableId" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/table/list' }" />
                </el-form-item> -->
                <el-form-item v-if="saveForm.data.sceneId" :label="t('auth.role.name.menuId')" prop="menuIdArr">
                    <my-cascader v-model="saveForm.data.menuIdArr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId } } }" :isPanel="true" :props="{ multiple: true }" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.sceneId" :label="t('auth.role.name.actionId')" prop="actionIdArr">
                    <my-transfer v-model="saveForm.data.actionIdArr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/action/list', param: { filter: { sceneId: saveForm.data.sceneId } } }" />
                </el-form-item>
                <el-form-item :label="t('auth.role.name.isStop')" prop="isStop">
                    <el-switch
                        v-model="saveForm.data.isStop"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
