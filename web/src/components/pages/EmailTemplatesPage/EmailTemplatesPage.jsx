import { useGetEmailTemplates } from "../../../api/notification";
import EmailTemplatesTemplate from "../../templates/EmailTemplatesTemplate/EmailTemplatesTemplate";

function EmailtTemplatesPage() {
  const { data: emailTemplatesData, isLoading: isGetEmailTemplatesLoading } =
    useGetEmailTemplates();

  return (
    <EmailTemplatesTemplate
      emailTemplatesData={emailTemplatesData}
      isGetEmailTemplatesLoading={isGetEmailTemplatesLoading}
    />
  );
}

export default EmailtTemplatesPage;
