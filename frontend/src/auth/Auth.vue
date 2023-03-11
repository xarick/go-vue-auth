<template>
    <n-space vertical>
        <n-layout>
            <n-layout has-sider>
                <Sidebar />
                <n-layout-content class="h-screen">
                    <Navbar />
                    <RouterView class="mx-10 my-5" />
                    <Footer />
                </n-layout-content>
            </n-layout>
        </n-layout>
    </n-space>
</template>
