export interface PaginationData {
  count: number
  total_count: number
  total_pages: number
  current_page: number
}

export interface PaginationProps {
  paginationData: PaginationData
  onPageChange: (page: number) => void
  loading?: boolean
}

export interface PaginationControlsProps {
  currentPage: number
  totalPages: number
  onPageChange: (page: number) => void
  loading?: boolean
}

export const calculatePageNumbers = (currentPage: number, totalPages: number): (number | string)[] => {
  if (totalPages <= 7) {
    return Array.from({ length: totalPages }, (_, i) => i + 1)
  }

  const pages: (number | string)[] = []
  
  if (currentPage <= 4) {
    for (let i = 1; i <= 5; i++) pages.push(i)
    pages.push('...')
    pages.push(totalPages)
  } else if (currentPage >= totalPages - 3) {
    pages.push(1)
    pages.push('...')
    for (let i = totalPages - 4; i <= totalPages; i++) pages.push(i)
  } else {
    pages.push(1)
    pages.push('...')
    for (let i = currentPage - 1; i <= currentPage + 1; i++) pages.push(i)
    pages.push('...')
    pages.push(totalPages)
  }
  
  return pages
}

export const formatPaginationInfo = (data: PaginationData): string => {
  const { count, total_count, current_page, total_pages } = data
  
  if (!count || count === 0) {
    return `แสดง 0 รายการ (หน้า ${current_page}/${total_pages})`
  }
  
  const startItem = ((current_page - 1) * count) + 1
  const endItem = Math.min(startItem + count - 1, total_count)
  
  return `แสดง ${startItem.toLocaleString()}-${endItem.toLocaleString()} จาก ${total_count.toLocaleString()} รายการ (หน้า ${current_page}/${total_pages})`
}

export const validatePageNumber = (page: number, totalPages: number): number => {
  if (page < 1) return 1
  if (page > totalPages) return totalPages
  return page
}