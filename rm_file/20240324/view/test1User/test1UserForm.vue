<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="user:" prop="user">
          <el-input v-model="formData.user" :clearable="true"  placeholder="请输入user" />
       </el-form-item>
        <el-form-item label="password:" prop="password">
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入password" />
       </el-form-item>
        <el-form-item label="刷步方式:" prop="pushType">
          <el-input v-model.number="formData.pushType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="今日步数:" prop="todayStepCnt">
          <el-input v-model.number="formData.todayStepCnt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTest1User,
  updateTest1User,
  findTest1User
} from '@/api/test1User'

defineOptions({
    name: 'Test1UserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            user: '',
            password: '',
            pushType: 0,
            todayStepCnt: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTest1User({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reTestUser
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTest1User(formData.value)
               break
             case 'update':
               res = await updateTest1User(formData.value)
               break
             default:
               res = await createTest1User(formData.value)
               break
           }
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
