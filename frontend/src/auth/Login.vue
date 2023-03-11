<script setup>
import LoginBgImage from "../components/img/bg-login.jpg";

const route = useRoute();

onMounted(async () => {
    if (route.query.code) {
        const data = (await window.axios.get("/oauth/connect", { params: { access_token: route.query.code } })).data;
        
        localStorage.setItem("user", JSON.stringify({ id: data.id, name: data.name, cbid: data.cbid, iabs_id: data.iabs_id }));

        $cookies.set("token_user", data.token_user.split("|")[1], "expiring time");

        window.location.href = "/cabinet/home";
    }
});
</script>
<template>
    <div :style="`background-image: url('${LoginBgImage}')`" class="bg-cover bg-center h-screen">
        <div class="bg-black/50 h-screen w-screen">
            <div class="flex items-center h-screen justify-between max-w-screen-md mx-auto">
                <div>
                    <img src="../components/img/logo.svg" class="w-72" />
                    <p class="text-white mt-1 uppercase text-xs tracking-widest text-right select-none">
                        Birgalikda yuksalish sari!
                    </p>
                </div>
                <div class="block p-10 backdrop-blur-sm rounded-lg shadow-xl relative bg-white max-w-sm text-center">
                    <p class="text-gray-800 font-bold text-lg uppercase whitespace-nowrap">
                        Tizim nomi
                    </p>
                    <div class="flex items-center justify-center mt-6">
                        <a href="/oauth/login" class="btn">InHouse ID bilan kirish</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
