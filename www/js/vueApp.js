let app = new Vue({
    el: '#app',
    data: {
        language: {
            login: 'usuário',
            password: 'senha',
            send: 'enviar'
        },
        user: '',
        password: ''
    },
    methods: {
        sendCredentials: function() {
            console.log(this.user);
            console.log(this.password);
        }
    }
});