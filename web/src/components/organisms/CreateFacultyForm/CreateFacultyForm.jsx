import { Stack, TextField, Button } from "@mui/material";
import { useTranslation } from "react-i18next";
import { useFormik } from "formik";
import { useCreateFaculty } from "../../../api/faculty";
import { getCreateFacultyValidationSchema } from "../../../helpers/validation";

function CreateFacultyForm() {
  const { t } = useTranslation();
  const { mutate: createFacultyMutate, isLoading: isCreateFacultyLoading } =
    useCreateFaculty();

  const formik = useFormik({
    initialValues: {
      name: "",
    },
    validationSchema: getCreateFacultyValidationSchema(),
    onSubmit: handleSubmit,
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack
        sx={{
          flexDirection: {
            sx: "column",
            md: "row",
          },
          gap: 2,
        }}
      >
        <TextField
          id="name"
          name="name"
          size="small"
          label={t("organisms.createFacultyForm.form.name.label")}
          value={formik.values.name}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.name && !!formik.errors.name}
          helperText={formik.touched.name && t(formik.errors.name)}
        />
        <Stack>
          <Button
            variant="contained"
            type="submit"
            disabled={isCreateFacultyLoading}
          >
            {t("organisms.createFacultyForm.form.button.submit")}
          </Button>
        </Stack>
      </Stack>
    </form>
  );

  function handleSubmit(values, actions) {
    createFacultyMutate(values);

    actions.resetForm();
  }
}

export default CreateFacultyForm;
