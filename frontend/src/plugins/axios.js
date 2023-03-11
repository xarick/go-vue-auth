import axios from 'axios'


$cookies.get('token_user') ? axios.defaults.headers.common['Authorization'] = 'Bearer ' + $cookies.get('token_user') : ''
// axios.defaults.baseURL = "/api/"

axios.interceptors.request.use(function (config) {
    return config;
}, function (error) {
    window.useMessage.error(error)
    return Promise.reject(error);
});

axios.interceptors.response.use((response) => responseSuccess(response), (error) => responseError(error));

const responseSuccess = (data) => {
    if (data.status === 200) {
        switch (data.config.type) {
            case 'create':
                window.useMessage.success("Muvaffaqiyatli yaratildi")
                break
            case 'update':
                window.useMessage.success("Muvaffaqiyatli yangilandi")
                break
            case 'destroy':
                window.useMessage.warning("Muvaffaqiyatli o'chirildi")
                break
            default:
                break
        }
    }
    return data
}

const responseError = (data) => {
    console.log(data.response);
    window.useMessage.error(data.response.data.message)
    switch (data.response.status) {
        case 500:
            // window.location.href = '/cabinet/error'
            break
        case 401:
            $cookies.remove("token_user")
            localStorage.removeItem('user')
            window.location.href = '/login'
            break
        default:
            break
    }
    // if (data.response.data.code) {
    //     window.useMessage.error("xatolik")
    //     switch (data.response.data.code) {
    //         case 403:
    //             window.location.href = '/cabinet/error?status=403'
    //             break
    //         case 500:
    //             window.location.href = '/cabinet/error'
    //             break
    //         default:
    //             break
    //     }
    // } else {
    //     console.log(data.response);
    //     switch (data.response.status) {
    //         case 500:
    //             // window.location.href = '/cabinet/error'
    //             window.useMessage.error(i18n.global.t('window.useMessage.error.server'))
    //             break
    //         case 401:
    //             $cookies.remove("token_user")
    //             localStorage.removeItem('user')
    //             window.location.href = '/login'
    //             break
    //         default:
    //             break
    //     }
    // }
}

export default axios
