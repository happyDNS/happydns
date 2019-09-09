import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import router from './router'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

require('./assets/css/glyphicon.css')

Vue.use(BootstrapVue)

Vue.config.productionTip = false

new Vue({
  router,
  render: function (h) { return h(App) }
}).$mount('#app')

Vue.filter('toto', function (input, filter) {
  console.log(input)
  return input
})
Vue.filter('nsclass', function (input) {
  switch (input) {
    case 1:
      return 'IN'
    case 3:
      return 'CH'
    case 4:
      return 'HS'
    case 254:
      return 'NONE'
    default:
      return '##'
  }
})
Vue.filter('nsrrtype', function (input) {
  switch (input) {
    case 1: return 'A'
    case 2: return 'NS'
    case 3: return 'MD'
    case 4: return 'MF'
    case 5: return 'CNAME'
    case 6: return 'SOA'
    case 7: return 'MB'
    case 8: return 'MG'
    case 9: return 'MR'
    case 10: return 'NULL'
    case 11: return 'WKS'
    case 12: return 'PTR'
    case 13: return 'HINFO'
    case 14: return 'MINFO'
    case 15: return 'MX'
    case 16: return 'TXT'
    case 17: return 'RP'
    case 18: return 'AFSDB'
    case 19: return 'X25'
    case 20: return 'ISDN'
    case 21: return 'RT'
    case 22: return 'NSAP'
    case 23: return 'NSAP-PTR'
    case 24: return 'SIG'
    case 25: return 'KEY'
    case 26: return 'PX'
    case 27: return 'GPOS'
    case 28: return 'AAAA'
    case 29: return 'LOC'
    case 30: return 'NXT'
    case 31: return 'EID'
    case 32: return 'NIMLOC'
    case 33: return 'SRV'
    case 34: return 'ATMA'
    case 35: return 'NAPTR'
    case 36: return 'KX'
    case 37: return 'CERT'
    case 38: return 'A6'
    case 39: return 'DNAME'
    case 40: return 'SINK'
    case 41: return 'OPT'
    case 42: return 'APL'
    case 43: return 'DS'
    case 44: return 'SSHFP'
    case 45: return 'IPSECKEY'
    case 46: return 'RRSIG'
    case 47: return 'NSEC'
    case 48: return 'DNSKEY'
    case 49: return 'DHCID'
    case 50: return 'NSEC3'
    case 51: return 'NSEC3PARAM'
    case 52: return 'TLSA'
    case 53: return 'SMIMEA'
    case 55: return 'HIP'
    case 56: return 'NINFO'
    case 57: return 'RKEY'
    case 58: return 'TALINK'
    case 59: return 'CDS'
    case 60: return 'CDNSKEY'
    case 61: return 'OPENPGPKEY'
    case 62: return 'CSYNC'
    case 63: return 'ZONEMD'
    case 99: return 'SPF'
    case 100: return 'UINFO'
    case 101: return 'UID'
    case 102: return 'GID'
    case 103: return 'UNSPEC'
    case 104: return 'NID'
    case 105: return 'L32'
    case 106: return 'L64'
    case 107: return 'LP'
    case 108: return 'EUI48'
    case 109: return 'EUI64'
    case 249: return 'TKEY'
    case 250: return 'TSIG'
    case 251: return 'IXFR'
    case 252: return 'AXFR'
    case 253: return 'MAILB'
    case 254: return 'MAILA'
    case 256: return 'URI'
    case 257: return 'CAA'
    case 258: return 'AVC'
    case 259: return 'DOA'
    case 260: return 'AMTRELAY'
    case 32768: return 'TA'
    case 32769: return 'DLV'
    default: return '#'
  }
})
