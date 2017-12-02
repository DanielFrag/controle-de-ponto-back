Vue.component('login', {
    props: ['language'],
    template: `<ul>
        <li>
            <input type="text" v-model="$root.user" :placeholder="language.login"/>
        </li>
        <li>
            <input type="password" v-model="$root.password" :placeholder="language.password"/>
        </li>
        <li>
            <button v-on:click="$root.sendCredentials">{{language.send}}</button>
        </li>
    </ul>`
});