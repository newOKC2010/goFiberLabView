import Loading from '@/components/loading/mainLoading';

export default function TableLoading() {
  return (
    <div className="py-16">
      <Loading 
        message="กำลังโหลดข้อมูล"
        size="md"
        fullScreen={false}
        delay={999999}
      />
    </div>
  );
}
