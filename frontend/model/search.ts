import type {IssueFullResponse} from "~/model/model";
import type {PullRequestResponse} from "~/model/pull";

export interface SearchResponse {
    type: String,
    score: Number,
    summary: String,
    content: IssueFullResponse | PullRequestResponse
}
