import os
from pdf2image import convert_from_path
from PIL import Image

def merge_images(images):
    images = [Image.open(x) for x in images]
    widths, heights = zip(*(i.size for i in images))

    total_width = max(widths) if len(widths) < 3 else max(widths) * 2
    max_height = max(heights) * 2 if len(heights) > 1 else max(heights)

    new_im = Image.new('RGB', (total_width, max_height), color='white')

    x_offset = 0
    y_offset = 0
    for i, im in enumerate(images):
        if i % 2 == 0 and i != 0:
            y_offset += max_height // 2
            x_offset = 0
        new_im.paste(im, (x_offset, y_offset))
        x_offset += im.size[0]

    return new_im

def convert_pdf_to_images(pdf_path):
    images = convert_from_path(pdf_path)
    image_paths = []
    for i, image in enumerate(images):
        image_path = f"{pdf_path[:-4]}_{i}.png"
        image.save(image_path, "PNG")
        image_paths.append(image_path)

    merged_image_paths = []
    for i in range(0, len(image_paths), 4):
        merged_image = merge_images(image_paths[i:i+4])
        merged_image_path = f"{pdf_path[:-4]}_merged_{i}.png"
        merged_image.save(merged_image_path, "PNG")
        merged_image_paths.append(merged_image_path)

    return merged_image_paths

pdf_path = "xxx.pdf"
convert_pdf_to_images(pdf_path)