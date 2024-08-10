export interface ReviewResponse {
    id: number;
    author: string;
    body: string;
    state: string;
    createdAt: string;  // ISO date string
    updatedAt: string;  // ISO date string
}

export interface ReviewCommentResponse {
    id: number;
    author: string;
    body: string;
    path: string;
    position: number;
    createdAt: string;  // ISO date string
    updatedAt: string;  // ISO date string
}

export interface PullRequestResponse {
    id: number;
    number: number;
    title: string;
    state: string;
    locked: boolean;
    draft: boolean;
    author: string;
    assignees: string[];
    labels: string[];
    createdAt: string;  // ISO date string
    updatedAt: string;  // ISO date string
    closedAt?: string;  // ISO date string, optional
    mergedAt?: string;  // ISO date string, optional
    milestone?: string;  // optional
    repository: string;
    body: string;
    mergeCommitSha?: string;  // optional
    headBranch: string;
    baseBranch: string;
    reviewList: ReviewResponse[];  // List of associated reviews
    reviewCommentList: ReviewCommentResponse[];  // List of associated review comments
}
