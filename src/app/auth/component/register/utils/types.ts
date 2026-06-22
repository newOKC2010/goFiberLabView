export interface RegisterRequest {
  cid: string;
  full_name: string;
  email: string;
  facility_type: string;
  facility_code: string;
  facility_name: string;
}

export interface RegisterErrors {
  cid: string;
  fullName: string;
  email: string;
  facilityType: string;
  facilityCode: string;
  facilityName: string;
}

export interface RegisterModalProps {
  isOpen: boolean;
  onClose: () => void;
}
