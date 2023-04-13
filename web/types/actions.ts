export interface Action {
  id: string;
  user: string;
  status: string;
  created_at: string;
}

export interface Incomplete_Action {
  id: string;
  action_id: string;
  user_id: string;
  request_id: string;
  request_type: string;
}

export interface Approve_Action {
  request_id: string;
  request_type: string;
  user_id: string;
}

export interface Reject_Action {
  request_id: string;
  user_id: string;
}
