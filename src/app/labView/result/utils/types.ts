export interface LabItem {
  lab_items_name: string;
  lab_order_result: string | null;
}

export interface LabGroupedResult {
  order_date: string;
  items: LabItem[];
}

export interface LabResultResponse {
  success: boolean;
  total: number;
  pt_name: string;
  results: LabGroupedResult[];
}
