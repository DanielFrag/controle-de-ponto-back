webpackJsonp([6],{211:function(t,e,n){var i=n(212);"string"==typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);n(25)("5e46cbf0",i,!0)},212:function(t,e,n){e=t.exports=n(24)(void 0),e.push([t.i,".custom-btn[data-v-4355ee39]{margin:10px}",""])},213:function(t,e,n){"use strict";var i=n(2);e.a={name:"cSearch",components:{QBtn:i.c,QDatetimeRange:i.f,QIcon:i.g,QInput:i.i},data:function(){return{range:{from:new Date,to:new Date}}},methods:{clickOnSubmitButton:function(){this.$refs.submitButton.click()},submit:function(){var t=this;if(!this.range.from||!this.range.to)return void i.t.create.negative({html:"Data em branco"});var e=new Date(this.range.from),n=new Date(this.range.to);e.setHours(0),e.setMinutes(0),n.setHours(23),n.setMinutes(59),this.$http.get("https://controledeponto.herokuapp.com/user/registers/"+e.getTime()+"/"+n.getTime(),{headers:{Authorization:"bearer "+this.$store.state.token}}).then(function(e){t.$store.commit("SET_TOKEN",e.body.token),e.body.dateRegisters&&e.body.dateRegisters.length?e.body.dateRegisters.sort(function(t,e){return new Date(t.timeStamp).getTime()-new Date(e.timeStamp).getTime()}):i.t.create.info({html:"Nenhum registro encontrado"}),t.$store.commit("ADD_REGISTERS",e.body.dateRegisters||[]),i.b.$emit("registers-stored")},function(e){4===parseInt(e.status/100)?t.$router.replace("/"):i.t.create.negative({html:"Erro ao consultar registros"})}).finally(function(){})}}}},214:function(t,e,n){"use strict";var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"layout-padding docs-input row justify-center"},[n("div",{staticStyle:{width:"500px","max-width":"90vw"}},[n("form",{on:{submit:function(e){e.preventDefault(),t.clickOnSubmitButton(e)}}},[n("p",{staticClass:"caption"},[t._v("Selecione os Dias")]),t._v(" "),n("q-datetime-range",{staticClass:"full-width",attrs:{color:"secondary"},model:{value:t.range,callback:function(e){t.range=e},expression:"range"}}),t._v(" "),n("q-btn",{ref:"submitButton",staticClass:"full-width custom-btn",attrs:{big:"",color:"primary"},on:{click:t.submit}},[t._v("\n        Enviar Consulta\n      ")])],1)])])},s=[],a={render:i,staticRenderFns:s};e.a=a},31:function(t,e,n){"use strict";function i(t){n(211)}Object.defineProperty(e,"__esModule",{value:!0});var s=n(213),a=n(214),r=n(3),o=i,c=r(s.a,a.a,!1,o,"data-v-4355ee39",null);e.default=c.exports}});