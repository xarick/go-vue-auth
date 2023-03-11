<script setup>
import { NIcon } from "naive-ui"
import { Shield, Users, Accessible, CheckupList, FileText, AlertTriangle, Home } from "@vicons/tabler"
import axios from "../plugins/axios"

const menus = ref([])

const icons = {
    Shield: Shield,
    Users: Users,
    Accessible: Accessible,
    FileText: FileText,
    CheckupList: CheckupList,
    AlertTriangle: AlertTriangle,
    Home: Home
}

const menuThemeOverrides = {
    borderRadius: "0.375rem",
    color: "#0000",
    groupTextColor: "rgb(118, 124, 130)",
    itemColorHover: "#e5e7ebFF",
    itemColorActive: "#e5e7ebFF",
    itemColorActiveHover: "#D1D5DBB3",
    itemColorActiveCollapsed: "#e5e7ebFF",
    itemTextColorActive: "#111827FF",
    itemTextColorActiveHover: "#111827FF",
    itemTextColorChildActive: "#111827FF",
    itemTextColorChildActiveHover: "#111827FF",
    itemIconColor: "#6b7280FF",
    itemIconColorHover: "#111827FF",
    itemIconColorActiveHover: "#111827FF",
    itemIconColorActive: "#111827FF",
    itemIconColorChildActive: "#111827FF",
    itemIconColorChildActiveHover: "#111827FF",
    itemIconColorCollapsed: "#6b7280FF",
    itemIconColorHorizontal: "#6b7280FF",
    arrowColor: "#6b7280FF",
    arrowColorHover: "#111827FF",
    arrowColorActive: "#111827FF",
    arrowColorActiveHover: "#111827FF",
    arrowColorChildActive: "#111827FF",
    arrowColorChildActiveHover: "#111827FF",
}

const inverted = ref(false)

const renderIcon = (icon) => {
    return () => h(NIcon, null, { default: () => h(icon) });
}

onMounted(async () => {
    const data = (await axios.get("/api/menu")).data

    menus.value = data.map((menu) => {
        return {
            label: menu.router && menu.children.length == 0 ? () => h(RouterLink, { to: { name: menu.router }, }, { default: () => menu.name }) : menu.name,
            key: menu.id,
            icon: renderIcon(icons[menu.icon]),
            children:
                menu.children.length > 0
                    ? menu.children.map((child) => {
                        return {
                            label: () => h(RouterLink, { to: { name: child.router }, }, { default: () => child.name }),
                            key: child.id,
                            icon: renderIcon(icons[child.icon]),
                        };
                    })
                    : null,
        };
    });
})
</script>


<template>
    <n-layout-sider bordered show-trigger collapse-mode="width" :collapsed-width="64" :width="240" :native-scrollbar="false"
        :inverted="inverted" class="h-screen lg:border-gray-200 bg-gray-100">
        <n-menu :inverted="inverted" :collapsed-width="64" :collapsed-icon-size="22" :options="menus"
            :theme-overrides="menuThemeOverrides" class="text-sm font-medium mt-6" />
    </n-layout-sider>
</template>
