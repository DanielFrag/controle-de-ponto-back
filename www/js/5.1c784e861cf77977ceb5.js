webpackJsonp([5,13],{174:function(t,e,a){"use strict";var o=a(2);e.a={name:"cInsert",components:{QBtn:o.c,QIcon:o.g,QInlineDatetime:o.h,QInput:o.i,QSelect:o.q},data:function(){return{date:new Date,obs:"",options:[{label:"",value:""},{label:"Entrada",value:"Entrada"},{label:"Saída",value:"Saída"}],select:""}},methods:{reset:function(){this.date=new Date,this.select="",this.obs=""},submitConfirm:function(){var t=this;o.a.create({title:"Confirmação",message:"Deseja realmente enviar o registro "+this.date.toLocaleString(),buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){t.date.setSeconds(0),t.date.setMilliseconds(0);var e="";t.select&&(e=""+t.select),t.obs&&(e=e+": "+t.obs),t.$http.post("https://controledeponto.herokuapp.com/user/register-date",{timeStamp:t.date,description:e},{headers:{Authorization:"bearer "+t.$store.state.token}}).then(function(e){t.$store.commit("SET_TOKEN",e.body.token),o.t.create.positive({html:"Registro enviado"})},function(e){4===parseInt(e.status/100)?t.$router.replace("/"):o.t.create.negative({html:"Erro ao criar registro"})}).finally(function(){t.reset()})}}]})}}}},175:function(t,e,a){"use strict";var o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"layout-padding docs-input row justify-center"},[a("form",[a("q-inline-datetime",{attrs:{type:"datetime",color:"primary",format24h:""},model:{value:t.date,callback:function(e){t.date=e},expression:"date"}}),t._v(" "),a("q-select",{attrs:{"stack-label":"Descrição (opcional)",inverted:"",color:"secondary",separator:"",options:t.options},model:{value:t.select,callback:function(e){t.select=e},expression:"select"}}),t._v(" "),a("q-input",{attrs:{type:"textarea",color:"secondary",inverted:"","min-rows":5,placeholder:"Observação (opcional)"},model:{value:t.obs,callback:function(e){t.obs=e},expression:"obs"}}),t._v(" "),a("br"),t._v(" "),a("br"),t._v(" "),a("br"),t._v(" "),a("q-btn",{staticClass:"fixed",staticStyle:{right:"18px",bottom:"18px"},attrs:{round:"",color:"primary",glossy:""},on:{click:t.submitConfirm}},[a("q-icon",{attrs:{name:"add"}})],1)],1)])},n=[],r={render:o,staticRenderFns:n};e.a=r},227:function(t,e,a){var o=a(228);"string"==typeof o&&(o=[[t.i,o,""]]),o.locals&&(t.exports=o.locals);a(25)("6661bca0",o,!0)},228:function(t,e,a){e=t.exports=a(24)(void 0),e.push([t.i,"",""])},229:function(t,e,a){"use strict";var o=a(2),n=a(28);e.a={name:"insert",components:{CInsert:n.default,QBtn:o.c,QIcon:o.g,QLayout:o.n,QToolbar:o.r,QToolbarTitle:o.s},methods:{back:function(){this.$router.go(-1)}}}},230:function(t,e,a){"use strict";var o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("q-layout",{ref:"layout",attrs:{view:"lHh Lpr fff","left-class":{"bg-grey-2":!0}}},[a("q-toolbar",[a("q-btn",{attrs:{flat:""},on:{click:t.back}},[a("q-icon",{attrs:{name:"keyboard_arrow_left"}})],1),t._v(" "),a("q-toolbar-title",[t._v("\n      Inserir Registro\n    ")])],1),t._v(" "),a("c-insert")],1)},n=[],r={render:o,staticRenderFns:n};e.a=r},28:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=a(174),n=a(175),r=a(3),s=r(o.a,n.a,!1,null,null,null);e.default=s.exports},35:function(t,e,a){"use strict";function o(t){a(227)}Object.defineProperty(e,"__esModule",{value:!0});var n=a(229),r=a(230),s=a(3),i=o,l=s(n.a,r.a,!1,i,"data-v-f71e871e",null);e.default=l.exports}});