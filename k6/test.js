// noinspection NpmUsedModulesInstalled

import http from 'k6/http';

export function setup() {
    // noinspection JSUnresolvedVariable
    console.log(`URL to be tested: ${__ENV.URL}`)
}

export const options = {
    discardResponseBodies: true,
    thresholds: {
        http_req_failed: ['rate<0.02'], // http errors should be less than 2%
    },
    scenarios: {
        default: {
            executor: 'ramping-arrival-rate',
            preAllocatedVUs: 16,
            maxVUs: 65536,
            stages: [
                {target: 100000, duration: '30s'},
                {target: 100000, duration: '10s'},
                {target: 0, duration: '30s'},
            ],
        },
    },
};

export default function () {
    const params = {
        timeout: "30s",
    };

    // noinspection JSUnresolvedVariable
    http.get(`${__ENV.URL}`, params);
}
