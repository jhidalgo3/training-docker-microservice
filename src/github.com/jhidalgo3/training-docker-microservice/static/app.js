const BaseUrl = getAbsolutePath();

function getAbsolutePath() {
  var pathArray = window.location.href.split( '/' );
  var protocol = pathArray[0];
  var host = pathArray[2];
  var url = protocol + '//' + host;
  return url  
}

function buildUrl (url) {
  console.log (BaseUrl)
  return  BaseUrl + url;
}

Vue.component('info-host', {
  props: ['config','info'],
  template: `
  <div>
    <section>
      <ul class="collection with-header">
        <li class="collection-header"><h4>Info</h4></li>
        <li class="collection-item"><pre>{{info}}</pre></li>
      </ul>
    </section>

      <section>
      <ul class="collection with-header">
        <li class="collection-header"><h4>Config</h4></li>
        <li class="collection-item"><pre>{{config}}</pre></li>
      </ul>
    </section>
 </div>
  `,
  computed: {
    
  }
});

const vm = new Vue({
  el: '#app',
  data: {
    results: {},
    config: '',
    info: '',
    loading: true,
    title: ''
  },
  mounted() {
    this.getConfig();
     this.getInfo()
  },
  methods: {
    getConfig() {
      this.loading = true;
      let url = buildUrl("/api/config");
      axios.get(url).then(response => {
        this.loading = false;
        console.log (response.data)
        this.results = response.data;
        this.config = jsyaml.dump(this.results)
      }).catch(error => { console.log(error); });
    },

    getInfo() {
      this.loading = true;
      let url = buildUrl("/api/info");
      axios.get(url).then(response => {
        this.loading = false;
        console.log (response.data)
        this.results = response.data;
        this.info = jsyaml.dump(this.results)
      }).catch(error => { console.log(error); });
    },

    update (){
        this.getConfig()
        this.getInfo()
    }
  }
});

