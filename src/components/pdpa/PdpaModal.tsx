'use client';

interface PdpaModalProps {
  onAcknowledge: () => void;
}

export default function PdpaModal({ onAcknowledge }: PdpaModalProps) {
  return (
    <div className="fixed inset-0 z-[100] bg-black/60 backdrop-blur-sm flex items-center justify-center px-4 py-8 animate-fade-in">
      <div className="bg-white rounded-2xl shadow-2xl w-full max-w-lg overflow-hidden">

        {/* Header */}
        <div className="bg-gradient-to-r from-blue-600 to-cyan-600 px-6 py-5">
          <div className="flex items-center gap-3">
            <span
              className="material-symbols-outlined text-white text-3xl"
              style={{ fontVariationSettings: "'wght' 300" }}
            >
              policy
            </span>
            <div>
              <h2 className="text-white font-bold text-lg">นโยบายคุ้มครองข้อมูลส่วนบุคคล</h2>
              <p className="text-blue-100 text-xs font-bold mt-0.5">
                พ.ร.บ. คุ้มครองข้อมูลส่วนบุคคล พ.ศ. 2562 (PDPA)
              </p>
            </div>
          </div>
        </div>

        {/* Body */}
        <div className="px-6 py-5 space-y-4 text-sm text-gray-700 max-h-[60vh] overflow-y-auto">

          <div className="bg-amber-50 border border-amber-200 rounded-xl px-4 py-3 flex gap-3">
            <span
              className="material-symbols-outlined text-amber-500 text-xl shrink-0 mt-0.5"
              style={{ fontVariationSettings: "'wght' 500" }}
            >
              warning
            </span>
            <p className="font-bold text-amber-800 leading-relaxed">
              ระบบนี้มีการประมวลผลข้อมูลสุขภาพซึ่งเป็น
              <span className="text-red-600"> ข้อมูลส่วนบุคคลที่มีความอ่อนไหว (Sensitive Personal Data)</span>
              &nbsp;ภายใต้การคุ้มครองของ พ.ร.บ. คุ้มครองข้อมูลส่วนบุคคล พ.ศ. 2562
            </p>
          </div>

          <Section icon="visibility" title="การบันทึกการเข้าถึงข้อมูล">
            การเข้าถึงข้อมูลผลการตรวจทางห้องปฏิบัติการของผู้ป่วยทุกครั้ง
            <span className="font-bold text-gray-800"> จะถูกบันทึกไว้ในระบบ (Audit Log)</span>{' '}
            ประกอบด้วย วันที่และเวลา ชื่อผู้ใช้งาน และข้อมูลที่ถูกเข้าถึง
            เพื่อให้สามารถตรวจสอบย้อนหลังได้ตามข้อกำหนดของกฎหมาย
          </Section>

          <Section icon="medical_information" title="วัตถุประสงค์การใช้ข้อมูล">
            ข้อมูลสุขภาพในระบบนี้จัดเก็บและเปิดเผยเพื่อ
            <span className="font-bold text-gray-800"> วัตถุประสงค์ด้านการดูแลรักษาสุขภาพเท่านั้น</span>{' '}
            ห้ามนำข้อมูลไปใช้เพื่อวัตถุประสงค์อื่น เผยแพร่ หรือส่งต่อโดยไม่ได้รับอนุญาต
          </Section>

          <Section icon="gavel" title="ความรับผิดชอบของผู้ใช้งาน">
            ผู้ใช้งานมีหน้าที่รักษาความลับของข้อมูลผู้ป่วย และต้องไม่กระทำการใด ๆ
            อันเป็นการละเมิดสิทธิ์ความเป็นส่วนตัวของผู้ป่วย
            การฝ่าฝืนอาจมีโทษตามกฎหมาย พ.ร.บ. คุ้มครองข้อมูลส่วนบุคคล
            และกฎหมายที่เกี่ยวข้อง
          </Section>

          <Section icon="lock" title="การรักษาความมั่นคงปลอดภัย">
            กรุณาออกจากระบบทุกครั้งเมื่อเสร็จสิ้นการใช้งาน
            และไม่อนุญาตให้ผู้อื่นใช้บัญชีของท่านในการเข้าถึงข้อมูล
          </Section>

          <p className="text-xs text-gray-400 font-bold border-t pt-3">
            หากท่านพบการใช้ข้อมูลที่ไม่ถูกต้อง กรุณาแจ้งเจ้าหน้าที่ผู้ดูแลระบบทันที
          </p>
        </div>

        {/* Footer */}
        <div className="px-6 py-4 bg-gray-50 border-t border-gray-100 flex flex-col sm:flex-row items-center justify-between gap-3">
          <p className="text-xs text-gray-400 font-bold text-center sm:text-left">
            การกดปุ่ม "รับทราบ" ถือว่าท่านได้อ่านและยอมรับ<br className="sm:hidden" />นโยบายข้างต้นแล้ว
          </p>
          <button
            onClick={onAcknowledge}
            className="flex items-center gap-2 px-6 py-2.5 bg-gradient-to-r from-blue-500 to-cyan-500 hover:from-blue-600 hover:to-cyan-600 text-white text-sm font-bold rounded-xl shadow transition cursor-pointer whitespace-nowrap"
          >
            <span
              className="material-symbols-outlined text-base"
              style={{ fontVariationSettings: "'wght' 700" }}
            >
              check_circle
            </span>
            รับทราบและเข้าใช้งาน
          </button>
        </div>
      </div>
    </div>
  );
}

function Section({ icon, title, children }: { icon: string; title: string; children: React.ReactNode }) {
  return (
    <div className="flex gap-3">
      <span
        className="material-symbols-outlined text-blue-500 text-lg shrink-0 mt-0.5"
        style={{ fontVariationSettings: "'wght' 400" }}
      >
        {icon}
      </span>
      <div>
        <p className="font-bold text-gray-800 mb-0.5">{title}</p>
        <p className="text-gray-600 leading-relaxed font-bold">{children}</p>
      </div>
    </div>
  );
}
