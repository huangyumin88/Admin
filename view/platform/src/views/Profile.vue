<script setup lang="tsx">
import md5 from 'js-md5'

const { t } = useI18n()
const adminStore = useAdminStore()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        nickname: adminStore.info.nickname,
        avatar: adminStore.info.avatar,
    } as { [propName: string]: any },
    rules: {
        account: [
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },
        ],
        phone: [
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },
        ],
        nickname: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        avatar: [
            { type: 'url', trigger: 'change', message: t('validation.upload') },
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
        ],
        password: [{ type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) }],
        repeatPassword: [
            { required: computed((): boolean => (saveForm.data.password ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
            {
                trigger: 'blur',
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password != saveForm.data.repeatPassword) {
                        callback(new Error())
                    }
                    callback()
                },
                message: t('validation.repeatPassword'),
            },
        ],
        passwordToCheck: [
            { required: computed((): boolean => (saveForm.data.account || saveForm.data.phone || saveForm.data.password ? true : false)), message: t('profile.tip.passwordToCheck') },
            { type: 'string', trigger: 'blur', min: 6, max: 30, message: t('validation.between.string', { min: 6, max: 30 }) },
            {
                trigger: 'blur',
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password && saveForm.data.password == saveForm.data.passwordToCheck) {
                        callback(new Error())
                    }
                    callback()
                },
                message: t('validation.newPasswordDiffOldPassword'),
            },
        ],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.account || delete param.account
            param.phone || delete param.phone
            param.password ? (param.password = md5(param.password)) : delete param.password
            delete param.repeatPassword
            param.passwordToCheck ? (param.passwordToCheck = md5(param.passwordToCheck)) : delete param.passwordToCheck
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/my/profile/update', param, true)
                //成功则更新用户信息
                for (let k in param) {
                    if (adminStore.info.hasOwnProperty(k)) {
                        adminStore.info[k] = param[k]
                    }
                }
            } catch (error) {}
            saveForm.loading = false
        })
    },
})
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
                <el-form-item :label="t('profile.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('profile.name.account')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.account', { account: adminStore.info.account ? adminStore.info.account : t('common.tip.notSet') })" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('profile.name.phone')" maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.phone', { phone: adminStore.info.phone ? adminStore.info.phone : t('common.tip.notSet') })" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('profile.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('profile.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('profile.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('profile.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.repeatPassword')" prop="repeatPassword">
                    <el-input v-model="saveForm.data.repeatPassword" :placeholder="t('profile.name.repeatPassword')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.passwordToCheck')" prop="passwordToCheck">
                    <el-input v-model="saveForm.data.passwordToCheck" :placeholder="t('profile.name.passwordToCheck')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.passwordToCheck')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
</template>
