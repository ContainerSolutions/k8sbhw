const url = 'http://localhost:8080/webapp/store';

Vue.component('RecordList', {
  props: ['records'],

  template: `
    <ul>
      <li v-for="record in records"> 
        <record :record="record" readonly></record>
      </li>
      <li>
        <record :record='{}'></record>
      </li>
    </ul>
  `
});

Vue.component('Record', {
  props: ['record', 'readonly'],

  template: `
    <div>
      <input v-model="record.Key" type="text" name="key" v-bind:readonly="readonly">
      <input v-model="record.Value" type="text" name="key">
      <button v-on:click="updateRecord(record)">v</button>
      <button v-on:click="deleteRecord(record)">x</button>
    </div>
  `,

  methods: {
    deleteRecord: function(record) {
      return axios.delete(url, {data: {Key: record.Key}})
        .then((response) => {
          vm.getRecords()
        }
      );
    },
    updateRecord: function(record) {
      return axios.put(url, {Key: record.Key, Value: record.Value})
        .then((response) => {
          vm.getRecords()
        }
      );
    }
  }
});

const vm = new Vue({
  el: '#app',
  data: {
    records: []
  },

  methods: {
    getRecords: function() {
      return axios.get(url)
        .then((response) => {
          this.records = response.data
      });
    },
  },

  mounted: function() {
    this.getRecords();
  }

});
