
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="邮箱:" prop="email">
    <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="第三方账号:" prop="thirdAccount">
    <el-input v-model="formData.thirdAccount" :clearable="true" placeholder="请输入第三方账号" />
</el-form-item>
        <el-form-item label="渠道uid:" prop="channelUid">
    <el-input v-model="formData.channelUid" :clearable="true" placeholder="请输入渠道uid" />
</el-form-item>
        <el-form-item label="密码:" prop="pwd">
    <el-input v-model="formData.pwd" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="用户昵称:" prop="nickname">
    <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入用户昵称" />
</el-form-item>
        <el-form-item label="头像:" prop="icon">
    <el-input v-model="formData.icon" :clearable="true" placeholder="请输入头像" />
</el-form-item>
        <el-form-item label="用户类型;0普通用户;1VIP:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入用户类型;0普通用户;1VIP" />
</el-form-item>
        <el-form-item label="vip用户的过期时间:" prop="vipExpiredAt">
    <el-input v-model.number="formData.vipExpiredAt" :clearable="true" placeholder="请输入vip用户的过期时间" />
</el-form-item>
        <el-form-item label="0-正常，1-锁定:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入0-正常，1-锁定" />
</el-form-item>
        <el-form-item label="注册来源:1邮箱注册;2 goole id;3 apple id:" prop="origin">
    <el-select v-model="formData.origin" clearable placeholder="请选择注册来源">
      <el-option :value="1" label="1 - 邮箱注册" />
      <el-option :value="2" label="2 - Google ID" />
      <el-option :value="3" label="3 - Apple ID" />
      <el-option :value="4" label="4" />
    </el-select>
</el-form-item>
        <el-form-item label="注册时候的系统语言:" prop="lan">
    <el-input v-model="formData.lan" :clearable="true" placeholder="请输入注册时候的系统语言" />
</el-form-item>
        <el-form-item label="转写语言:" prop="transLan">
    <el-input v-model="formData.transLan" :clearable="true" placeholder="请输入转写语言" />
</el-form-item>
        <el-form-item label="注册国家/地区编码:" prop="country">
    <el-input v-model="formData.country" :clearable="true" placeholder="请输入注册国家/地区编码" />
</el-form-item>
        <el-form-item label="用户所属行业:" prop="industry">
    <el-input v-model.number="formData.industry" :clearable="true" placeholder="请输入用户所属行业" />
</el-form-item>
        <el-form-item label="设备ID:" prop="deviceId">
    <el-input v-model="formData.deviceId" :clearable="true" placeholder="请输入设备ID" />
</el-form-item>
        <el-form-item label="注册ip:" prop="ip">
    <el-input v-model="formData.ip" :clearable="true" placeholder="请输入注册ip" />
</el-form-item>
        <el-form-item label="用户迁移状态:1迁移;2被迁移:" prop="syncStatus">
    <el-select v-model="formData.syncStatus" clearable placeholder="请选择迁移状态">
      <el-option :value="1" label="1 - 迁移" />
      <el-option :value="2" label="2 - 被迁移" />
    </el-select>
</el-form-item>
        <el-form-item label="迁移的用户uid:" prop="syncUid">
    <el-input v-model.number="formData.syncUid" :clearable="true" placeholder="请输入迁移的用户uid" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="删除写入时间戳:" prop="deletedAt">
    <el-input v-model.number="formData.deletedAt" :clearable="true" placeholder="请输入删除写入时间戳" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWavenoteUser,
  updateWavenoteUser,
  findWavenoteUser
} from '@/api/system/wavenoteUser'

defineOptions({
    name: 'WavenoteUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            email: '',
            thirdAccount: '',
            channelUid: '',
            pwd: '',
            nickname: '',
            icon: '',
            userType: undefined,
            vipExpiredAt: undefined,
            status: undefined,
            origin: undefined,
            lan: '',
            transLan: '',
            country: '',
            industry: undefined,
            deviceId: '',
            ip: '',
            syncStatus: undefined,
            syncUid: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWavenoteUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createWavenoteUser(formData.value)
               break
             case 'update':
               res = await updateWavenoteUser(formData.value)
               break
             default:
               res = await createWavenoteUser(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
