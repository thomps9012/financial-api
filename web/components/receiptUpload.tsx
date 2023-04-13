import ImageUploading, {
  ImageListType,
  ImageType,
} from "react-images-uploading";
import Image from "next/image";
import { useState } from "react";
let imageState: ImageType[];
export default function ReceiptUpload({ receipts, setReceipts }: any) {
  const format_image_list = () => {
    const image_state = receipts.map((receipt: string) => ({
      data_url: receipt,
      file: "An Image",
    }));
    return image_state;
  };
  const [imageList, setImageList] = useState(format_image_list());
  const onChange = (receiptList: ImageListType) => {
    const receiptArr: string[] = receiptList.map(
      (receipt: any) => receipt.data_url
    );
    console.log(receiptList);
    setImageList(receiptList);
    setReceipts(receiptArr);
  };
  const onError = ({ errors, files }: any) => {
    console.log("Error", errors, files);
  };
  const maxNumber = 5;
  return (
    <ImageUploading
      multiple
      value={imageList}
      onChange={onChange}
      maxNumber={maxNumber}
      onError={onError}
      acceptType={["png", "jpg", "pdf"]}
      dataURLKey="data_url"
    >
      {({
        imageList,
        onImageRemoveAll,
        onImageRemove,
        onImageUpload,
        isDragging,
        dragProps,
        errors,
      }) => (
        <div className="upload-image-wrapper">
          {errors && (
            <div className="error-container">
              {errors.maxNumber && (
                <span className="error-alert">
                  {"You've reached the image upload limit"}
                </span>
              )}
              {errors.acceptType && (
                <span className="error-alert">
                  Your attempting to upload a forbidden file type
                </span>
              )}
              {errors.maxFileSize && (
                <span className="error-alert">
                  Your file exceeds the max size
                </span>
              )}
              {errors.resolution && (
                <span className="error-alert">
                  Your file is not the correct resolution
                </span>
              )}
            </div>
          )}
          <br />
          <div
            onClick={onImageUpload}
            className="upload-area"
            style={
              isDragging
                ? { background: "cadetblue", opacity: "50%" }
                : undefined
            }
            {...dragProps}
          >
            <h2 className="description">
              Click or Drag and Drop <br /> to <br /> Upload Receipt Images
            </h2>
          </div>
          <br />
          {imageList.length > 0 && (
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
              }}
            >
              <a onClick={onImageRemoveAll} className="reject-btn">
                Remove All Receipts
              </a>
              <p style={{ color: "rgb(160, 95, 95)" }} className="req-overview">
                or Click an Image to Remove It
              </p>
            </div>
          )}
          <div className="image-container">
            {imageList.map((image, index) => (
              <div key={index} className="image-item">
                <a onClick={() => onImageRemove(index)}>
                  <Image
                    src={image["data_url"] || image}
                    alt=""
                    width="200"
                    height="200"
                  />
                </a>
              </div>
            ))}
          </div>
          <p style={{ textAlign: "right" }}>
            {receipts.length} Attached Receipt{receipts.length != 1 && "s"}
          </p>
        </div>
      )}
    </ImageUploading>
  );
}
