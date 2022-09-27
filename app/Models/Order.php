<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Order extends Model
{
    use HasFactory;
    protected $fillable = ['address', 'latitude', 'longitude', 'status'];

    public function products(){
        return $this->belongsToMany(Product::class)->using(OrderItem::class)->withPivot('quantity', 'price');
    }

    public function user(){
        return $this->belongsTo(User::class);
    }
}