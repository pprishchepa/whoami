// noinspection NpmUsedModulesInstalled

import {check} from 'k6';
import http from 'k6/http';

export function setup() {
    // noinspection JSUnresolvedVariable
    console.log(`URL to be tested: ${__ENV.URL}`)
}

export const options = {
    discardResponseBodies: true,
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
    const res = http.get(`${__ENV.URL}`, params);

    check(res, {
        'is status 200': (r) => r.status === 200,
    });
}
