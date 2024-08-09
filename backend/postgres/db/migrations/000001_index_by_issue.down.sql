-- Drop index on comments.issue_url
DROP INDEX IF EXISTS public.idx_comments_issue_url;

-- Drop index on issues.number
DROP INDEX IF EXISTS public.idx_issues_number;
