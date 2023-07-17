import http from 'k6/http'

    export let options = {
        vus: 1000,
        duration: '5s'
    }

    export default () => {
    http.get('http://host.docker.internal:8000/products')
 }