<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class OrderRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, mixed>
     */
    public function rules()
    {
        return [
            'user_id' => 'required',
            'address' => 'required', 
            'latitude' => 'required', 
            'longitude' => 'required', 
            'status' => 'required',
            'products' => 'required|array',
            'product.*.id' => 'required',
            'product.*.pivot.quantity' => 'required',
            'product.*.pivot.price' => 'required'
        ];
    }
}
