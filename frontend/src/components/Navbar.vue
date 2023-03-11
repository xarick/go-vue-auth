<script setup>
import { NIcon } from "naive-ui";
import { PersonCircleOutline as UserIcon, LogOutOutline, LogOut } from "@vicons/ionicons5";
import { GridDots } from "@vicons/tabler";

const renderIcon = (icon) => {
    return () => h(NIcon, null, { default: () => h(icon) });
}

const details = reactive({
    inverted: false,
    show: {
        modal: false,
        dropDown: false
    },
    user: JSON.parse(localStorage.getItem('user'))
})

const options = [
    {
        label: "Profil",
        key: "profile",
        icon: renderIcon(UserIcon)
    },
    {
        label: "Chiqish",
        key: "logout",
        icon: renderIcon(LogOutOutline)
    },
    {
        label: "Chiqish (ID)",
        key: "logout-with-sso",
        icon: renderIcon(LogOut)
    },
]

const logout = (key) => {
    if (key == 'logout') {
        $cookies.remove("token_user")
        localStorage.removeItem('user')
        window.location.href = '/logout'
    } else if (key == 'profile') {
        details.show.modal = true
    }
    else {
        $cookies.remove("token_user")
        localStorage.removeItem('user')
        window.location.href = '/logout?withSSO=true'
    }
}
</script>

<template>
    <n-layout-header :inverted="details.inverted" bordered class="h-14 items-center flex justify-between px-4">
        <p class="text-lg font-bold">Asosiy</p>
        <div class="items-center justify-center flex">
            <n-popover trigger="hover" placement="bottom" :show-arrow="false">
                <template #trigger>
                    <n-button class="rounded-full flex items-center justify-center p-2 h-10 w-10 mx-2" quaternary>
                        <GridDots class="w-5" />
                    </n-button>
                </template>
                <div class="grid gap-4 grid-cols-3 justify-items-center place-content-center">
                    <div
                        class="rounded-md hover:bg-slate-100 p-3 w-[70px] h-[70px] flex flex-col items-center justify-center cursor-pointer">
                        <img src="../components/img/inhouse.png" class="w-9" alt="inhouse">
                        <small class="font-bold">Inhouse</small>
                    </div>
                    <div
                        class="rounded-md hover:bg-slate-100 p-3 w-[70px] h-[70px] flex flex-col items-center justify-center cursor-pointer">
                        <img src="../components/img/inhouse.png" class="w-9" alt="inhouse">
                        <small class="font-bold">Inhouse</small>
                    </div>
                    <div
                        class="rounded-md hover:bg-slate-100 p-3 w-[70px] h-[70px] flex flex-col items-center justify-center cursor-pointer">
                        <img src="../components/img/inhouse.png" class="w-9" alt="inhouse">
                        <small class="font-bold">Inhouse</small>
                    </div>
                </div>
            </n-popover>
            <n-dropdown trigger="click" :options="options" @select="logout">
                <n-button
                    class="rounded-full h-10 w-10 flex items-center uppercase leading-none p-2 font-bold justify-center"
                    @click="details.show.dropDown = true">
                    jt
                </n-button>
            </n-dropdown>
        </div>
    </n-layout-header>

    <n-modal v-model:show="details.show.modal" preset="dialog">
        <template #header>
            <div>Xodim haqida</div>
        </template>
        <div class="mt-4 flex row-auto">
            <div
                class="col-auto bg-green-500 rounded-full w-16 h-16 items-center justify-center font-bold text-xl flex text-white uppercase">
                jt
            </div>
            <div class="col-auto ml-3">
                <span class="font-bold uppercase">F.I.O:</span>
                {{ details.user.name }}
                <br>
                <span class="font-bold uppercase">cbid:</span>
                {{ details.user.cbid }}
                <br>
                <span class="font-bold uppercase">iabs id:</span>
                {{ details.user.iabs_id }}
            </div>
        </div>
    </n-modal>
</template>