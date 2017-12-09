webpackJsonp([2,8,12],{168:function(t,e,a){var r=a(169);"string"==typeof r&&(r=[[t.i,r,""]]),r.locals&&(t.exports=r.locals);a(25)("092c9ce4",r,!0)},169:function(t,e,a){e=t.exports=a(24)(void 0),e.push([t.i,".custom-btn[data-v-21c7b629]{margin:10px}",""])},170:function(t,e,a){"use strict";var r=a(2);e.a={name:"cDelete",components:{QBtn:r.c,QDatetime:r.e,QIcon:r.g,QInput:r.i},data:function(){return{day:new Date}},methods:{submit:function(){var t=this;if(!this.day)return void r.t.create.negative({html:"Data em branco"});var e=new Date(this.day.getFullYear(),this.day.getMonth(),this.day.getDate(),0,0,0),a=new Date(this.day.getFullYear(),this.day.getMonth(),this.day.getDate(),23,59,59);this.$http.get("https://controledeponto.herokuapp.com/user/registers/"+e.getTime()+"/"+a.getTime(),{headers:{Authorization:"bearer "+this.$store.state.token}}).then(function(e){t.$store.commit("SET_TOKEN",e.body.token),e.body.dateRegisters&&e.body.dateRegisters.length&&(e.body.dateRegisters.sort(function(t,e){return new Date(t.timeStamp).getTime()-new Date(e.timeStamp).getTime()}),t.$store.commit("ADD_REGISTERS",e.body.dateRegisters),r.b.$emit("registers-stored"))},function(e){4===parseInt(e.status/100)?t.$router.replace("/"):r.t.create.negative({html:"Erro ao consultar registros"})}).finally(function(){})}}}},171:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"layout-padding docs-input row justify-center"},[a("div",{staticStyle:{width:"500px","max-width":"90vw"}},[a("form",[a("p",{staticClass:"caption"},[t._v("Selecione o Dia")]),t._v(" "),a("q-datetime",{staticClass:"full-width",attrs:{color:"secondary"},model:{value:t.day,callback:function(e){t.day=e},expression:"day"}}),t._v(" "),a("q-btn",{staticClass:"full-width custom-btn",attrs:{big:"",color:"primary"},on:{click:t.submit}},[t._v("\n        Buscar Registros\n      ")])],1)])])},n=[],i={render:r,staticRenderFns:n};e.a=i},172:function(t,e,a){"use strict";var r=a(2);e.a={name:"cRegisters",components:{QItem:r.j,QItemMain:r.k,QItemSeparator:r.l,QItemTile:r.m,QList:r.o},computed:{registers:function(){return this.$store.state.registers}},created:function(){this.$q.events.$on("registers-stored",this.populateView)},beforeDestroy:function(){this.$q.events.$off("registers-stored",this.populateView)},data:function(){return{data:[]}},methods:{populateView:function(){var t=this;this.data=[],this.registers.forEach(function(e){var a=new Date(e.timeStamp);t.data.push({date:a.toLocaleDateString(),time:a.toLocaleTimeString(),description:e.description,id:e._id})})},deleteConfirm:function(t){var e=this;r.a.create({title:"Confirmação",message:"Deseja realmente deletar o registro",buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){e.$http.delete("https://controledeponto.herokuapp.com/user/remove-register/"+t,{headers:{Authorization:"bearer "+e.$store.state.token}}).then(function(t){e.$store.commit("SET_TOKEN",t.body.token),e.data=[],r.t.create.positive({html:"Item deletado com sucesso"})},function(t){4===parseInt(t.status/100)?e.$router.replace("/"):r.t.create.negative({html:"Erro ao deletar registro"})}).finally(function(){})}}]})}}}},173:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",t._l(t.data,function(e){return a("q-list",{key:e.id,attrs:{"inset-separator":""}},[a("q-item",{on:{click:function(a){t.deleteConfirm(e.id)}}},[a("q-item-main",[a("q-item-tile",{attrs:{label:""}},[t._v("\n          "+t._s(e.time)+"\n        ")]),t._v(" "),a("q-item-tile",{attrs:{sublabel:""}},[t._v("\n          "+t._s(e.description)+"\n        ")])],1)],1)],1)}))},n=[],i={render:r,staticRenderFns:n};e.a=i},219:function(t,e,a){var r=a(220);"string"==typeof r&&(r=[[t.i,r,""]]),r.locals&&(t.exports=r.locals);a(25)("39314db4",r,!0)},220:function(t,e,a){e=t.exports=a(24)(void 0),e.push([t.i,"",""])},221:function(t,e,a){"use strict";var r=a(2),n=a(26),i=a(27);e.a={name:"delete",components:{QBtn:r.c,CDelete:n.default,CRegisters:i.default,QIcon:r.g,QLayout:r.n,QToolbar:r.r,QToolbarTitle:r.s},methods:{back:function(){this.$router.go(-1)}}}},222:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("q-layout",{ref:"layout",attrs:{view:"lHh Lpr fff","left-class":{"bg-grey-2":!0}}},[a("q-toolbar",[a("q-btn",{attrs:{flat:""},on:{click:t.back}},[a("q-icon",{attrs:{name:"keyboard_arrow_left"}})],1),t._v(" "),a("q-toolbar-title",[t._v("\n      Deletar Registro\n    ")])],1),t._v(" "),a("c-delete"),t._v(" "),a("c-registers")],1)},n=[],i={render:r,staticRenderFns:n};e.a=i},26:function(t,e,a){"use strict";function r(t){a(168)}Object.defineProperty(e,"__esModule",{value:!0});var n=a(170),i=a(171),o=a(3),s=r,c=o(n.a,i.a,!1,s,"data-v-21c7b629",null);e.default=c.exports},27:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=a(172),n=a(173),i=a(3),o=i(r.a,n.a,!1,null,null,null);e.default=o.exports},33:function(t,e,a){"use strict";function r(t){a(219)}Object.defineProperty(e,"__esModule",{value:!0});var n=a(221),i=a(222),o=a(3),s=r,c=o(n.a,i.a,!1,s,"data-v-de5d9328",null);e.default=c.exports}});