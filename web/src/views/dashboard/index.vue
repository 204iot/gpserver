<template>
  <div>
    <el-amap vid="amap" class="amap-demo" zoom="2" :center="center" :events="events">
      <el-amap-marker
        v-for="(m, index) in markers"
        :key="index"
        :vid="index"
        :position="m.position"
        :label="m.label"
      />
    </el-amap>

    <div class="toolbar">
      position: [{{ lng }}, {{ lat }}] address: {{ address }}
    </div>
  </div>
</template>

<script>

import { devices } from '@/api/device'

export default {
  data() {
    const self = this
    return {
      timer: null,
      center: [121.59996, 31.197646],
      lng: 0,
      lat: 0,
      address: '',
      loaded: false,
      markers: [
        {
          position: [121.5273285, 31.21515044],
          label: { content: 'cc', offset: [10, -10] }
        }
      ],
      events: {
        click(e) {
          const { lng, lat } = e.lnglat
          self.lng = lng
          self.lat = lat
          // 这里通过高德 SDK 完成。
          var geocoder = new AMap.Geocoder({
            radius: 1000,
            extensions: 'all'
          })
          geocoder.getAddress([lng, lat], function(status, result) {
            if (status === 'complete' && result.info === 'OK') {
              if (result && result.regeocode) {
                self.address = result.regeocode.formattedAddress
                self.$nextTick()
              }
            }
          })
        }
      }
    }
  },
  mounted() {
    this.timer = setInterval(this.getDevices, 2000)
  },
  beforeDestroy() {
    clearInterval(this.timer)
  },
  methods: {
    getDevices() {
      devices().then(response => {
        const markers = []
        for (let i = 0, len = response.length; i < len; i++) {
          markers.push({
            position: [parseFloat(response[i].longitude), parseFloat(response[i].latitude)],
            // title: '测试',
            label: { content: response[i].id, offset: [10, -10] }
          })
        }
        console.log(markers)
        this.markers = markers
      })
    }
  }
}
</script>

<style scoped>
  .amap-demo {
    height: 500px;
  }
</style>
