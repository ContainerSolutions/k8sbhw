var url = 'http://localhost:8080/webapp/store';

vm = new Vue({
  el: '#app',
  data: {
    items: [ ]
  },
  methods:{
    getData: function() {
      fetch(url).then(function(response){
        return response.json();
      }).then(function(data) {
        vm.items = data;
      }).catch(function() {
        console.log("Fetch failed: " + url);
      });
    }
  },
  mounted: function() {
    this.getData();
  }
});

Vue.component('item')
