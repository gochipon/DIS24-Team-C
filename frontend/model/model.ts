export interface CommentResponse {
    id: number;
    author: string;
    body: string;
    createdAt: string;  // ISO date string
    updatedAt: string;  // ISO date string
}
export interface User {
    id: number;
    url: string;
    type: string;
    login: string;
    node_id: string;
    html_url: string;
    gists_url: string;
    repos_url: string;
    avatar_url: string;
    events_url: string;
    site_admin: boolean;
    gravatar_id: string;
    starred_url: string;
    followers_url: string;
    following_url: string;
    organizations_url: string;
    subscriptions_url: string;
    received_events_url: string;
}

export interface Draft {
    Bool: boolean;
    Valid: boolean;
}

export interface Reactions {
    "+1": number;
    "-1": number;
    url: string;
    eyes: number;
    heart: number;
    laugh: number;
    hooray: number;
    rocket: number;
    confused: number;
    total_count: number;
}

export interface PullRequest {
    url: string;
    diff_url: string;
    html_url: string;
    merged_at: string;
    patch_url: string;
}

export interface IssueResponse {
    id: number;
    url: string;
    body: string;
    user: User;
    draft: Draft;
    state: string;
    title: string;
    labels: any[]; // Assuming labels are an array, replace with appropriate type if needed
    locked: boolean;
    number: number;
    node_id: string;
    user_id: {
        Int64: number;
        Valid: boolean;
    };
    assignee: null | string;
    comments: number;
    html_url: string;
    assignees: string; // If assignees should be an array, change this type accordingly
    closed_at: string;
    milestone: null | string;
    reactions: Reactions;
    created_at: string;
    events_url: string;
    labels_url: string;
    repository: string;
    updated_at: string;
    comments_url: string;
    pull_request: PullRequest;
    state_reason: null | string;
    timeline_url: string;
    repository_url: string;
    active_lock_reason: null | string;
    author_association: string;
    performed_via_github_app: null | string;
    _airbyte_raw_id: string;
    _airbyte_extracted_at: string;
    _airbyte_generation_id: number;
    _airbyte_meta: {
        changes: any[]; // Assuming changes is an array, replace with appropriate type if needed
        sync_id: number;
    };
}


export interface IssueFullResponse {
    issue: IssueResponse;
    comments: CommentResponse[];
}
