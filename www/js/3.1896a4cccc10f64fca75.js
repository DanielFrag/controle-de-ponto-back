webpackJsonp([3,9],{215:function(t,e,a){var n=a(216);"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);a(25)("5439e209",n,!0)},216:function(t,e,a){e=t.exports=a(24)(void 0),e.push([t.i,".settings-list[data-v-30ba8dbc]{width:500px;max-width:90vw}",""])},217:function(t,e,a){"use strict";var n=a(2);e.a={name:"cInsert",components:{QBtn:n.c,QCollapsible:n.d,QInput:n.i,QList:n.o},data:function(){return{oldPass:"",newPass:""}},methods:{clickOnSubmitButton:function(){this.$refs.submitButton.click()},submitConfirm:function(t,e){var a=this;n.a.create({title:"Confirmação",message:"Deseja realmente mudar sua senha?",buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){a.$http.put("https://controledeponto.herokuapp.com/change-password",{oldPassword:a.oldPass,newPassword:a.newPass},{headers:{Authorization:"bearer "+a.$store.state.token}}).then(function(t){a.$store.commit("SET_TOKEN",t.body.token),n.t.create.positive({html:"Senha alterada"})},function(t){4===parseInt(t.status/100)?a.$router.replace("/"):n.t.create.negative({html:"Erro ao mudar senha"})}).finally(function(){e()})}}]})}}}},218:function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"layout-padding docs-input row justify-center"},[a("q-list",{staticClass:"settings-list"},[a("q-collapsible",{attrs:{group:"settings",label:"Mudar senha"}},[a("form",{on:{submit:function(e){e.preventDefault(),t.clickOnSubmitButton(e)}}},[a("q-input",{attrs:{type:"password",placeholder:"Senha atual"},model:{value:t.oldPass,callback:function(e){t.oldPass=e},expression:"oldPass"}}),t._v(" "),a("q-input",{attrs:{type:"password",placeholder:"Nova senha"},model:{value:t.newPass,callback:function(e){t.newPass=e},expression:"newPass"}}),t._v(" "),a("q-btn",{ref:"submitButton",attrs:{loader:"",big:"",color:"primary"},on:{click:t.submitConfirm}},[t._v("\n          Enviar\n          "),a("span",{attrs:{slot:"loading"},slot:"loading"},[t._v("Enviando...")])])],1)])],1)],1)},s=[],o={render:n,staticRenderFns:s};e.a=o},241:function(t,e,a){var n=a(242);"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);a(25)("0932cdab",n,!0)},242:function(t,e,a){e=t.exports=a(24)(void 0),e.push([t.i,"",""])},243:function(t,e,a){"use strict";var n=a(2),s=a(32);e.a={name:"delete",components:{CChangePassword:s.default,QBtn:n.c,QIcon:n.g,QLayout:n.n,QToolbar:n.r,QToolbarTitle:n.s},methods:{back:function(){this.$router.go(-1)}}}},244:function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("q-layout",{ref:"layout",attrs:{view:"lHh Lpr fff","left-class":{"bg-grey-2":!0}}},[a("q-toolbar",[a("q-btn",{attrs:{flat:""},on:{click:t.back}},[a("q-icon",{attrs:{name:"keyboard_arrow_left"}})],1),t._v(" "),a("q-toolbar-title",[t._v("\n      Configurações\n    ")])],1),t._v(" "),a("c-change-password")],1)},s=[],o={render:n,staticRenderFns:s};e.a=o},32:function(t,e,a){"use strict";function n(t){a(215)}Object.defineProperty(e,"__esModule",{value:!0});var s=a(217),o=a(218),r=a(3),i=n,l=r(s.a,o.a,!1,i,"data-v-30ba8dbc",null);e.default=l.exports},39:function(t,e,a){"use strict";function n(t){a(241)}Object.defineProperty(e,"__esModule",{value:!0});var s=a(243),o=a(244),r=a(3),i=n,l=r(s.a,o.a,!1,i,"data-v-4e7960cc",null);e.default=l.exports}});