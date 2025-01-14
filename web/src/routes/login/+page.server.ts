import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport';
import type { LayoutLoad } from '../$types';
import { ReviewServiceClient } from '../../proto/review_tool.client';

export const load: LayoutLoad = async () => {
    var transport = new TwirpFetchTransport({
        baseUrl: "https://crev.abhayf.com:8080",
    });

    var client = new ReviewServiceClient(transport);
    client.sayHello({name: "world"}).then((response) => {
        console.log(response);
    });

	return {};
};

