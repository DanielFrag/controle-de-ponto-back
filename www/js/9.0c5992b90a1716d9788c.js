webpackJsonp([9],{215:function(t,a,e){var n=e(216);"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);e(25)("5439e209",n,!0)},216:function(t,a,e){a=t.exports=e(24)(void 0),a.push([t.i,".settings-list[data-v-30ba8dbc]{width:500px;max-width:90vw}",""])},217:function(t,a,e){"use strict";var n=e(2);a.a={name:"cInsert",components:{QBtn:n.c,QCollapsible:n.d,QInput:n.i,QList:n.o},data:function(){return{oldPass:"",newPass:""}},methods:{clickOnSubmitButton:function(){this.$refs.submitButton.click()},submitConfirm:function(t,a){var e=this;n.a.create({title:"Confirmação",message:"Deseja realmente mudar sua senha?",buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){e.$http.put("https://controledeponto.herokuapp.com/change-password",{oldPassword:e.oldPass,newPassword:e.newPass},{headers:{Authorization:"bearer "+e.$store.state.token}}).then(function(t){e.$store.commit("SET_TOKEN",t.body.token),n.t.create.positive({html:"Senha alterada"})},function(t){4===parseInt(t.status/100)?e.$router.replace("/"):n.t.create.negative({html:"Erro ao mudar senha"})}).finally(function(){a()})}}]})}}}},218:function(t,a,e){"use strict";var n=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"layout-padding docs-input row justify-center"},[e("q-list",{staticClass:"settings-list"},[e("q-collapsible",{attrs:{group:"settings",label:"Mudar senha"}},[e("form",{on:{submit:function(a){a.preventDefault(),t.clickOnSubmitButton(a)}}},[e("q-input",{attrs:{type:"password",placeholder:"Senha atual"},model:{value:t.oldPass,callback:function(a){t.oldPass=a},expression:"oldPass"}}),t._v(" "),e("q-input",{attrs:{type:"password",placeholder:"Nova senha"},model:{value:t.newPass,callback:function(a){t.newPass=a},expression:"newPass"}}),t._v(" "),e("q-btn",{ref:"submitButton",attrs:{loader:"",big:"",color:"primary"},on:{click:t.submitConfirm}},[t._v("\n          Enviar\n          "),e("span",{attrs:{slot:"loading"},slot:"loading"},[t._v("Enviando...")])])],1)])],1)],1)},s=[],o={render:n,staticRenderFns:s};a.a=o},32:function(t,a,e){"use strict";function n(t){e(215)}Object.defineProperty(a,"__esModule",{value:!0});var s=e(217),o=e(218),i=e(3),r=n,l=i(s.a,o.a,!1,r,"data-v-30ba8dbc",null);a.default=l.exports}});