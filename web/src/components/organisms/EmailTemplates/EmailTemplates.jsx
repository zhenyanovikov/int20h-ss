import { Stack } from "@mui/material";
import EmailTemplate from "../../molecules/EmailTemplate/EmailTemplate";

function EmailTemplates({ emailTemplates }) {
  return (
    <Stack spacing={4}>
      {emailTemplates.map((emailTemplate) => (
        <EmailTemplate key={emailTemplate.id} emailTemplate={emailTemplate} />
      ))}
    </Stack>
  );
}

export default EmailTemplates;
