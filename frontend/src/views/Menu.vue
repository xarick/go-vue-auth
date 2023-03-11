<script setup>
import { Shield, Pencil, Plus } from "@vicons/tabler"
import axios from "axios";

const menus = ref([]);
const details = reactive({
    modal: {
        style: {
            width: "600px"
        },
        title: "Menyu yangilash",
        show: false
    },
    type: 'main',
    mains: [],
    body: {}
})
const getMenus = onMounted(async (type = null) => {
    details.type = type
    menus.value = (await axios.get('/api/menu', { params: { type: details.type } })).data
    const data = (await axios.get('/api/menu')).data
    details.mains = []
    data.map(item => {
        details.mains.push({
            label: item.name,
            value: item.id
        })
    })
})
const create = async () => {
    if (await axios.post('/api/menu', details.body, { type: 'create' })) {
        getMenus()
        details.modal.show = false
    }
}
const show = async (id) => {
    details.body = (await axios.get('/api/menu/' + id)).data
    details.modal.title = "Menyuni yangilash"
    details.modal.show = true
}
const update = async () => {
    if (await axios.put('/api/menu/' + details.body.id, details.body, { type: 'update' })) {
        getMenus()
        details.body = {}
        details.modal.title = "Menyu qo'shish"
        details.modal.show = false
    }
}

</script>

<template>
    <div>
        <div class="grid grid-flow-col mb-6">
            <div class="flex justify-start">
                <n-button-group horizontal>
                    <n-button type="info" @click="getMenus('main')" ghost>
                        Asosiy menyu
                    </n-button>
                    <n-button type="info" @click="getMenus('sub')" ghost>
                        Sub menyu
                    </n-button>
                </n-button-group>
            </div>
            <div class="flex justify-end">
                <n-tooltip placement="left" trigger="hover">
                    <template #trigger>
                        <n-button type="primary" @click="details.modal.show = true" ghost>
                            <Plus class="w-4" />
                        </n-button>
                    </template>
                    Yangi qo'shish
                </n-tooltip>
            </div>
        </div>
        <div class="grid grid-cols-6 gap-4">
            <button v-for="menu in menus" @click="show(menu.id)" v-show="menu.id != 1 && menu.id != 2"
                class="group relative border rounded-md p-4 shadow-blue-400/20 hover:shadow-xl hover:shadow-blue-400/20 transition-shadow duration-300 text-center">
                <div class="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                    <Pencil class="w-6 text-slate-500" />
                </div>
                <div class="flex items-center justify-center bg-blue-500 rounded-full w-9 h-9 p-2 mx-auto">
                    <Shield class="w-8 text-white" />
                </div>
                <div class="uppercase mt-4 font-bold text-xs">{{ menu.name }}</div>
            </button>
        </div>

        <n-modal v-model:show="details.modal.show" class="custom-card" preset="card" :style="details.modal.style"
            :title="details.modal.title" :bordered="false" size="huge">
            <div class="items-center justify-center">
                <span class="mb-3">Nomi(lotin)</span>
                <n-input v-model:value="details.body.name_oz" type="text" clearable />
            </div>
            <div class="items-center justify-center my-4">
                <span class="mb-3">Nomi(kirill)</span>
                <n-input v-model:value="details.body.name_uz" type="text" clearable />
            </div>
            <div class="items-center justify-center">
                <span class="mb-3">Nomi(русский)</span>
                <n-input v-model:value="details.body.name_ru" type="text" clearable />
            </div>
            <div v-if="details.type == 'sub'" class="items-center justify-center mt-4">
                <span class="mb-3">Katta menyu</span>
                <n-select v-model:value="details.body.parent_id" :options="details.mains" />
            </div>
            <n-button type="primary" @click="details.body.id ? update() : create()" class="mt-4" ghost>
                Saqlash
            </n-button>
        </n-modal>
    </div>
</template>