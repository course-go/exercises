import { check } from 'k6';
import { Httpx } from 'https://jslib.k6.io/httpx/0.1.0/index.js';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import { sleep } from 'k6';

export const options = {
    scenarios: {
        data: {
            executor: 'constant-vus',
            exec: 'data',
            vus: 3,
            duration: '1h',
            gracefulStop: '1h',
        },
    },
};

const session = new Httpx({
    baseURL: __ENV.URL || "http://localhost:8080",
    timeout: 1000, // 1s
})

export function data() {
    createData()
}

function createData() {
    const count = randomIntBetween(0, 128)
    const url = `/data?count=${count}`
    const res = session.post(url, null, null);
    console.log(res.url)
    check(res, {
        'is status 200': (r) => r.status === 200,
    });
    sleep(0.1)
}
